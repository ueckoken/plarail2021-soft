const remoteVideo = document.getElementById("remote_video");
const roomIdInput = document.getElementById("room_id");
remoteVideo.controls = true;
class MoMoConnecter {
  constructor() {
    this.peerConnection = null;
    this.dataChannel = null;
    this.candidates = [];
    this.hasReceivedSdp = false;
    this.stream = null;
    // iceServer を定義
    this.iceServers = [{ urls: "stun:stun.l.google.com:19302" }];
    // peer connection の 設定
    this.peerConnectionConfig = {
      iceServers: this.iceServers,
    };
    console.log("書き換え");

    const wsUrl = "ws://127.0.0.1:8080/ws";
    this.ws = new WebSocket(wsUrl);
    this.ws.onopen = this.onWsOpen.bind(this);
    this.ws.onerror = this.onWsError.bind(this);
    this.ws.onmessage = this.onWsMessage.bind(this);
    this.prepareNewConnection = this.prepareNewConnection.bind(this);
    this.connect = this.connect.bind(this);
    this.disconnect = this.disconnect.bind(this);
    this.drainCandidate = this.drainCandidate.bind(this);
    this.addIceCandidate = this.addIceCandidate.bind(this);
    this.sendIceCandidate = this.sendIceCandidate.bind(this);
    this.sendSdp = this.sendSdp.bind(this);
    this.makeAnswer = this.makeAnswer.bind(this);
    this.makeOffer = this.makeOffer.bind(this);
    this.setAnswer = this.setAnswer.bind(this);
    this.setOffer = this.setOffer.bind(this);
    this.getStream = this.getStream.bind(this);
  }
  getStream() {
    return this.stream;
  }
  onWsError(error) {
    console.error("ws onerror() ERROR:", error);
  }
  onWsOpen(event) {
    console.log("ws open()");
  }
  onWsMessage(event) {
    console.log("ws onmessage() data:", event.data);
    const message = JSON.parse(event.data);
    if (message.type === "offer") {
      console.log("Received offer ...");
      const offer = new RTCSessionDescription(message);
      console.log("offer: ", offer);
      this.setOffer(offer);
    } else if (message.type === "answer") {
      console.log("Received answer ...");
      const answer = new RTCSessionDescription(message);
      console.log("answer: ", answer);
      this.setAnswer(answer);
    } else if (message.type === "candidate") {
      console.log("Received ICE candidate ...");
      const candidate = new RTCIceCandidate(message.ice);
      console.log("candidate: ", candidate);
      if (this.hasReceivedSdp) {
        this.addIceCandidate(candidate);
      } else {
        this.candidates.push(candidate);
      }
    } else if (message.type === "close") {
      console.log("peer connection is closed ...");
    }
  }
  connect() {
    console.group();
    if (!this.peerConnection) {
      console.log("make Offer");
      this.makeOffer();
    } else {
      console.warn("peer connection already exists.");
    }
    console.groupEnd();
  }
  disconnect() {
    console.group();
    if (this.peerConnection) {
      if (this.peerConnection.iceConnectionState !== "closed") {
        this.peerConnection.close();
        this.peerConnection = null;
        if (this.ws && this.ws.readyState === 1) {
          const message = JSON.stringify({ type: "close" });
          this.ws.send(message);
        }
        console.log("sending close message");
        cleanupVideoElement(remoteVideo);
        return;
      }
    }
    console.log("peerConnection is closed.");
    console.groupEnd();
  }
  drainCandidate() {
    this.hasReceivedSdp = true;
    this.candidates.forEach((candidate) => {
      this.addIceCandidate(candidate);
    });
    this.candidates = [];
  }
  addIceCandidate(candidate) {
    if (this.peerConnection) {
      this.peerConnection.addIceCandidate(candidate);
    } else {
      console.error("PeerConnection does not exist!");
    }
  }
  sendIceCandidate(candidate) {
    console.log("---sending ICE candidate ---");
    const message = JSON.stringify({ type: "candidate", ice: candidate });
    console.log("sending candidate=" + message);
    this.ws.send(message);
  }
  prepareNewConnection() {
    const peer = new RTCPeerConnection(this.peerConnectionConfig);
    this.dataChannel = peer.createDataChannel("serial");
    if ("ontrack" in peer) {
      if (isSafari()) {
        let tracks = [];
        peer.ontrack = (event) => {
          console.log("-- peer.ontrack()");
          tracks.push(event.track);
          // safari で動作させるために、ontrack が発火するたびに MediaStream を作成する
          let mediaStream = new MediaStream(tracks);
          playVideo(remoteVideo, mediaStream);
          this.stream = mediaStream;
        };
      } else {
        let mediaStream = new MediaStream();
        playVideo(remoteVideo, mediaStream);
        peer.ontrack = (event) => {
          console.log("-- peer.ontrack()");
          mediaStream.addTrack(event.track);
        };
        this.stream = mediaStream;
      }
    } else {
      peer.onaddstream = (event) => {
        console.log("-- peer.onaddstream()");
        playVideo(remoteVideo, event.stream);
        this.stream = event.stream;
      };
    }

    peer.onicecandidate = (event) => {
      console.log("-- peer.onicecandidate()");
      if (event.candidate) {
        console.log(event.candidate);
        this.sendIceCandidate(event.candidate);
      } else {
        console.log("empty ice event");
      }
    };

    peer.oniceconnectionstatechange = () => {
      console.log("-- peer.oniceconnectionstatechange()");
      console.log(
        "ICE connection Status has changed to " + peer.iceConnectionState
      );
      switch (peer.iceConnectionState) {
        case "closed":
        case "failed":
        case "disconnected":
          break;
      }
    };
    peer.addTransceiver("video", { direction: "recvonly" });
    peer.addTransceiver("audio", { direction: "recvonly" });

    this.dataChannel.onmessage = function (event) {
      console.log(
        "Got Data Channel Message:",
        new TextDecoder().decode(event.data)
      );
    };

    return peer;
  }
  sendSdp(sessionDescription) {
    console.log("---sending sdp ---");
    const message = JSON.stringify(sessionDescription);
    console.log("sending SDP=" + message);
    this.ws.send(message);
  }
  async makeOffer() {
    this.peerConnection = this.prepareNewConnection();
    try {
      const sessionDescription = await this.peerConnection.createOffer({
        offerToReceiveAudio: true,
        offerToReceiveVideo: true,
      });
      console.log(
        "createOffer() success in promise, SDP=",
        sessionDescription.sdp
      );
      switch (document.getElementById("codec").value) {
        case "H264":
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "VP8");
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "VP9");
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "AV1");
          break;
        case "VP8":
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "H264");
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "VP9");
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "AV1");
          break;
        case "VP9":
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "H264");
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "VP8");
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "AV1");
          break;
        case "AV1":
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "H264");
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "VP8");
          sessionDescription.sdp = removeCodec(sessionDescription.sdp, "VP9");
          break;
      }
      await this.peerConnection.setLocalDescription(sessionDescription);
      console.log("setLocalDescription() success in promise");
      this.sendSdp(this.peerConnection.localDescription);
    } catch (error) {
      console.error("makeOffer() ERROR:", error);
    }
  }
  async makeAnswer() {
    console.log("sending Answer. Creating remote session description...");
    if (!this.peerConnection) {
      console.error("peerConnection DOES NOT exist!");
      return;
    }
    try {
      const sessionDescription = await this.peerConnection.createAnswer();
      console.log("createAnswer() success in promise");
      await this.peerConnection.setLocalDescription(sessionDescription);
      console.log("setLocalDescription() success in promise");
      this.sendSdp(this.peerConnection.localDescription);
      this.drainCandidate();
    } catch (error) {
      console.error("makeAnswer() ERROR:", error);
    }
  }

  // offer sdp を生成する
  async setOffer(sessionDescription) {
    if (this.peerConnection) {
      console.error("peerConnection already exists!");
    }
    this.peerConnection = prepareNewConnection();
    this.peerConnection.onnegotiationneeded = async function () {
      try {
        await this.peerConnection.setRemoteDescription(sessionDescription);
        console.log("setRemoteDescription(offer) success in promise");
        this.makeAnswer();
      } catch (error) {
        console.error("setRemoteDescription(offer) ERROR: ", error);
      }
    };
  }

  async setAnswer(sessionDescription) {
    if (!this.peerConnection) {
      console.error("peerConnection DOES NOT exist!");
      return;
    }
    try {
      await this.peerConnection.setRemoteDescription(sessionDescription);
      console.log("setRemoteDescription(answer) success in promise");
      this.drainCandidate();
    } catch (error) {
      console.error("setRemoteDescription(answer) ERROR: ", error);
    }
  }
}

function browser() {
  const ua = window.navigator.userAgent.toLocaleLowerCase();
  if (ua.indexOf("edge") !== -1) {
    return "edge";
  } else if (ua.indexOf("chrome") !== -1 && ua.indexOf("edge") === -1) {
    return "chrome";
  } else if (ua.indexOf("safari") !== -1 && ua.indexOf("chrome") === -1) {
    return "safari";
  } else if (ua.indexOf("opera") !== -1) {
    return "opera";
  } else if (ua.indexOf("firefox") !== -1) {
    return "firefox";
  }
  return;
}

function playVideo(element, stream) {
  element.srcObject = stream;
}
function isSafari() {
  return browser() === "safari";
}

function cleanupVideoElement(element) {
  element.pause();
  element.srcObject = null;
}

function play() {
  remoteVideo.play();
}
function removeCodec(orgsdp, codec) {
  const internalFunc = (sdp) => {
    const codecre = new RegExp("(a=rtpmap:(\\d*) " + codec + "/90000\\r\\n)");
    const rtpmaps = sdp.match(codecre);
    if (rtpmaps == null || rtpmaps.length <= 2) {
      return sdp;
    }
    const rtpmap = rtpmaps[2];
    let modsdp = sdp.replace(codecre, "");

    const rtcpre = new RegExp("(a=rtcp-fb:" + rtpmap + ".*\r\n)", "g");
    modsdp = modsdp.replace(rtcpre, "");

    const fmtpre = new RegExp("(a=fmtp:" + rtpmap + ".*\r\n)", "g");
    modsdp = modsdp.replace(fmtpre, "");

    const aptpre = new RegExp("(a=fmtp:(\\d*) apt=" + rtpmap + "\\r\\n)");
    const aptmaps = modsdp.match(aptpre);
    let fmtpmap = "";
    if (aptmaps != null && aptmaps.length >= 3) {
      fmtpmap = aptmaps[2];
      modsdp = modsdp.replace(aptpre, "");

      const rtppre = new RegExp("(a=rtpmap:" + fmtpmap + ".*\r\n)", "g");
      modsdp = modsdp.replace(rtppre, "");
    }

    let videore = /(m=video.*\r\n)/;
    const videolines = modsdp.match(videore);
    if (videolines != null) {
      //If many m=video are found in SDP, this program doesn't work.
      let videoline = videolines[0].substring(0, videolines[0].length - 2);
      const videoelems = videoline.split(" ");
      let modvideoline = videoelems[0];
      videoelems.forEach((videoelem, index) => {
        if (index === 0) return;
        if (videoelem == rtpmap || videoelem == fmtpmap) {
          return;
        }
        modvideoline += " " + videoelem;
      });
      modvideoline += "\r\n";
      modsdp = modsdp.replace(videore, modvideoline);
    }
    return internalFunc(modsdp);
  };
  return internalFunc(orgsdp);
}
const momoConnecter = new MoMoConnecter();
window.connectMomo = momoConnecter.connect;
window.disconnectMomo = momoConnecter.disconnect;
const SkywayPeer = require("skyway-js");
const webSocket = new WebSocket("wss://127.0.0.1:8081/");

let skywayPeer = null;
let roomId = null;

window.connectReceiver = () => {
  skywayPeer = new SkywayPeer({
    key: "c07e8954-ce1b-4783-a45e-e8421ece83ce",
    debug: 3,
  });
  roomId = roomIdInput.value;
  skywayPeer.on("open", () => {
    webSocket.send(
      JSON.stringify({
        msg_type: "connect_sender",
        peer_id: skywayPeer.id,
        room_id: roomIdInput.value,
      })
    );
  });

  skywayPeer.on("call", (mediaConnection) => {
    const stream = momoConnecter.getStream();
    console.log(stream);
    console.log("on call");
    mediaConnection.answer(stream);
  });
};
window.disconnectReceiver = () => {
  skywayPeer.destroy();
  skywayPeer = null;
  webSocket.send(
    JSON.stringify({
      msg_type: "exit_room",
      room_id: roomId,
    })
  );
};

const cameraStream = null;
window.openCamera = () => {
  cameraStream = await getDeviceStream({ video: true, audio: false });
};
window.closeCamera = () => {
  if (cameraStream) {
    cameraStream.destroy();
  }
};
