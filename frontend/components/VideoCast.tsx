import { string } from "fp-ts";
import { NullType } from "io-ts";
import React, { FC, useEffect, useState, useRef } from "react";
import Peer, { MediaConnection, SfuRoom, MeshRoom } from "skyway-js";
interface PeerIdProp {
  peerId: string;
}
type MediaStreamWithPeerId = MediaStream & PeerIdProp;
interface Prop {}
const SW_WSURL = "wss://webrtc.chofufes2021.gotti.dev/";
const SKYWAY_APIKEY =
  process.env.SKYWAY_APIKEY === undefined
    ? "c07e8954-ce1b-4783-a45e-e8421ece83ce"
    : process.env.SKYWAY_APIKEY;
const SKYWAY_DEBUG_LEVEL = 2;
/*
window.closeVideoConnection = (roomId) => {
  const room = rooms[roomId];
  stopVideo(room["video_element"]);
  room["skyway_room"].close();
  delete rooms[roomId];

  webSocket.send(
    JSON.stringify({
      msg_type: "exit_room",
      room_id: roomId,
    })
  );
};*/
interface Room {
  skyway_room: SfuRoom | MeshRoom | null;
  //video_ref: HTMLAnchorElement | null;
  room_id: string | null;
}
const VideoCast: FC<Prop> = ({}) => {
  const [castingStream, setCastingStream] = useState<MediaStream | null>(null);
  const [webSocket, setWebSocket] = useState<WebSocket>();
  const [skywayPeer, setSkywayPeer] = useState<Peer>();
  const [rooms, setRooms] = useState<{ [room_id: string]: Room }>({});
  const videoRef = useRef<HTMLVideoElement>(null);
  useEffect(() => {
    const ws = new WebSocket(SW_WSURL);
    const peer = new Peer({
      key: SKYWAY_APIKEY,
      debug: SKYWAY_DEBUG_LEVEL,
    });
    setWebSocket(ws);
    setSkywayPeer(peer);
    let sendFuncs: Array<Function> = [];

    let isConnectWebSocket = false;
    let isConnectSkywayPeer = false;
    let roomId = "aaa";
    const room: Room = { skyway_room: null, room_id: roomId };
    const nextRooms = {
      ...rooms,
      [roomId]: room,
    };
    setRooms(nextRooms);
    const sendFunc = () => {
      ws.send(
        JSON.stringify({
          msg_type: "connect_receiver",
          room_id: roomId,
        })
      );
    };
    if (ws.readyState == ws.OPEN) {
      sendFunc();
    } else {
      sendFuncs.push(sendFunc);
    }
    ws.addEventListener("open", (event) => {
      isConnectWebSocket = true;
      if (!isConnectSkywayPeer) {
        return;
      }
      for (const func of sendFuncs) {
        func();
      }
      sendFuncs = [];
    });

    peer.on("open", () => {
      isConnectSkywayPeer = true;
      if (!isConnectWebSocket) {
        return;
      }
      for (const func of sendFuncs) {
        func();
      }
      sendFuncs = [];
    });
    ws.addEventListener("message", (event) => {
      const message = JSON.parse(event.data);
      console.log(message);
      const peerId = message["peer_id"];
      const roomId = message["room_id"];
      const skywayRoomId = message["skyway_room_id"];

      console.log("joinroom");
      const skywayRoom = peer.joinRoom(skywayRoomId, {
        mode: "sfu",
      });
      const room: Room = { room_id: roomId, skyway_room: skywayRoom };
      const nextRooms = {
        ...rooms,
        [roomId]: room,
      };
      if (skywayRoom) {
        skywayRoom.on("stream", (stream: MediaStreamWithPeerId) => {
          const streamPeerId = stream.peerId;
          console.log("on stream");
          if (streamPeerId == peerId) {
            setCastingStream(stream);
          }
        });
      }
      setRooms(nextRooms);
    });
    return () => {
      ws.close();
    };
  }, []);

  useEffect(() => {
    if (videoRef && videoRef.current) {
      if ("srcObject" in videoRef.current) {
        videoRef.current.srcObject = castingStream;
      } else {
        //videoRef.current.src = castingStream; // window.URL.createObjectURL(castingStream);
      }
      videoRef.current.play();
      videoRef.current.volume = 0;
    }
  }, [castingStream]);
  return (
    <React.Fragment>
      <video
        width={400}
        height={400}
        ref={videoRef}
        autoPlay
        playsInline
      ></video>
    </React.Fragment>
  );
};

export default VideoCast;
