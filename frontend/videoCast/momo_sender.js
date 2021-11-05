const Ayame = require("@open-ayame/ayame-web-sdk");
const remoteVideo = document.getElementById("remote_video");
const localVideo = document.getElementById("local_video");
const roomIdInput = document.getElementById("room_id");
async function getDeviceStream(option) {
  console.log("wrap navigator.getUserMadia with Promise");
  return new Promise(function (resolve, reject) {
    navigator.getUserMedia(option, resolve, reject);
  });
}
remoteVideo.controls = true;
localVideo.controls = true;

const signalingUrl = "wss://ayame-labo.shiguredo.jp/signaling";

const options = Ayame.defaultOptions;
options.signalingKey = "wyL5nvVEuX09H3qrLVNjGQq1gshPF1WZ60rCLaoohYEyQoRv";

let conn = null;
window.disconnectMomo = () => {
  if (conn) {
    conn.disconnect();
  }
};
let videoCodec = null;
window.onChangeVideoCodec = () => {
  videoCodec = document.getElementById("video_codec").value;
  if (videoCodec == "none") {
    videoCodec = null;
  }
};
window.connectMomo = async () => {
  options.video.codec = videoCodec;
  console.log(options.video.codec);
  conn = Ayame.connection(signalingUrl, "ruu413@ayame-momo", options, true);
  await conn.connect(null);
  conn.on("open", ({ authzMetadata }) => console.log(authzMetadata));
  conn.on("disconnect", (e) => {
    stopVideo(remoteVideo);
  });
  conn.on("addstream", (e) => {
    console.log("addstream");
    playVideo(remoteVideo, e.stream);
  });
};

function playVideo(element, stream) {
  element.srcObject = stream;
}
function stopVideo(element) {
  element.pause();
  element.srcObject = null;
}
const SkywayPeer = require("skyway-js");
const webSocket = new WebSocket(SW_WSURL);

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
    const selector = document.getElementById("select_source");
    const selectedIdx = selector.selectedIndex;
    const selected = selector.options[selectedIdx].value;
    console.log(selected);
    if (selected == "momo") {
      stream = momoConnecter.getStream();
    } else {
      stream = cameraStream;
    }
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

let cameraStream = null;
window.connectCamera = async () => {
  cameraStream = await getDeviceStream({ video: true, audio: false });
  playVideo(localVideo, cameraStream);
};
window.disconnectCamera = () => {
  if (cameraStream) {
    cameraStream.destroy();
  }
  stopVideo(localVideo);
};
