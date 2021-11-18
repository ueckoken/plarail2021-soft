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
options.signalingKey = AYAME_SIGNALING_KEY;

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
let momoStream = null;
window.connectMomo = async () => {
  options.video.codec = videoCodec;
  console.log(options.video.codec);
  conn = Ayame.connection(signalingUrl, "ruu413@ayame-momo", options, true);
  await conn.connect(null);
  conn.on("open", ({ authzMetadata }) => console.log(authzMetadata));
  conn.on("disconnect", (e) => {
    momoStream = null;
    stopVideo(remoteVideo);
  });
  conn.on("addstream", (e) => {
    console.log("addstream");
    momoStream = e.stream;
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
let skywayRoom = null;
window.connectReceiver = () => {
  skywayPeer = new SkywayPeer({
    key: SKYWAY_APIKEY,
    debug: SKYWAY_DEBUG_LEVEL,
  });
  roomId = roomIdInput.value;
  skywayPeer.on("open", () => {
    const selector = document.getElementById("select_source");
    const selectedIdx = selector.selectedIndex;
    const selected = selector.options[selectedIdx].value;
    let stream = null;
    if (selected == "momo") {
      stream = momoStream;
    } else {
      stream = cameraStream;
    }
    skywayRoom = skywayPeer.joinRoom(roomId, {
      mode: "sfu",
      stream: stream,
    });
    skywayRoom.on("open", () => {
      webSocket.send(
        JSON.stringify({
          msg_type: "connect_sender",
          peer_id: skywayPeer.id,
          room_id: roomId,
          sender_token: SENDER_TOKEN,
        })
      );
    });
  });
};
window.disconnectReceiver = () => {
  skywayPeer.destroy();
  skywayPeer = null;
  skywayRoom = null;
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
