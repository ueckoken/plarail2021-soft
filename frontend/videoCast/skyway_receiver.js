const Peer = require("skyway-js");
navigator.getUserMedia =
  navigator.getUserMedia ||
  navigator.webkitGetUserMedia ||
  navigator.mozGetUserMedia ||
  navigator.msGetUserMedia;

RTCPeerConnection =
  window.RTCPeerConnection ||
  window.webkitRTCPeerConnection ||
  window.mozRTCPeerConnection;
RTCSessionDescription =
  window.RTCSessionDescription ||
  window.webkitRTCSessionDescription ||
  window.mozRTCSessionDescription;

function playVideo(element, stream) {
  if ("srcObject" in element) {
    element.srcObject = stream;
  } else {
    element.src = window.URL.createObjectURL(stream);
  }
  element.play();
  element.volume = 0;
}

function stopVideo(element) {
  if ("srcObject" in element) {
    element.srcObject = null;
  } else {
    element.src = null;
  }
}
const webSocket = new WebSocket(SW_WSURL);
let sendFuncs = [];
const rooms = {};

webSocket.onopen = (event) => {
  for (const func of sendFuncs) {
    func();
  }
  sendFuncs = [];
};

webSocket.onmessage = (event) => {
  const message = JSON.parse(event.data);
  console.log(message);
  const peerId = message["peer_id"];
  const roomId = message["room_id"];
  const room = rooms[roomId];
  room["peer"] = new Peer({
    key: SKYWAY_APIKEY,
    debug: SKYWAY_DEBUG_LEVEL,
  });
  room["peer"].on("open", () => {
    console.log("joinroom");
    //room["media_connection"] = room["peer"].call(peerId);
    room["skyway_room"] = room["peer"].joinRoom(roomId, {
      mode: "sfu",
    });
    room["skyway_room"].on("stream", (stream) => {
      const streamPeerId = stream.peerId;
      console.log("on stream");
      if (streamPeerId == peerId) {
        playVideo(room["video_element"], stream);
      }
    });
  });
};

window.openVideoConnection = (id, roomId) => {
  rooms[roomId] = {
    video_element: document.getElementById(id),
    peer: null,
  };
  const sendFunc = () => {
    webSocket.send(
      JSON.stringify({
        msg_type: "connect_receiver",
        room_id: roomId,
      })
    );
  };
  if (webSocket.readyState == webSocket.OPEN) {
    sendFunc();
  } else {
    sendFuncs.push(sendFunc);
  }
};
window.closeVideoConnection = (roomId) => {
  const room = rooms[roomId];
  stopVideo(room["video_element"]);
  room["peer"].destroy();
  delete rooms[roomId];

  webSocket.send(
    JSON.stringify({
      msg_type: "exit_room",
      room_id: roomId,
    })
  );
};
