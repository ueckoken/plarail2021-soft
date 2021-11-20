import { string } from "fp-ts"
import { NullType } from "io-ts"
import React, { FC, useEffect, useState, useRef } from "react"
import Peer, { MediaConnection, SfuRoom, MeshRoom } from "skyway-js"
interface PeerIdProp {
  peerId: string
}
type MediaStreamWithPeerId = MediaStream & PeerIdProp
interface Prop {
  roomIds: string[]
}
const SW_WSURL = "wss://webrtc.chofufes2021.gotti.dev/"
const SKYWAY_APIKEY =
  process.env.SKYWAY_APIKEY === undefined
    ? "c07e8954-ce1b-4783-a45e-e8421ece83ce"
    : process.env.SKYWAY_APIKEY
const SKYWAY_DEBUG_LEVEL = 2
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
  skyway_room: SfuRoom | MeshRoom | null
  //video_ref: HTMLAnchorElement | null;
  room_id: string | null
}
const VideoCast: FC<Prop> = ({ roomIds }) => {
  const [castingStreams, setCastingStreams] = useState<{
    [room_id: string]: MediaStream | null
  }>({})
  const [webSocket, setWebSocket] = useState<WebSocket>()
  const [skywayPeer, setSkywayPeer] = useState<Peer>()
  const [rooms, setRooms] = useState<{ [room_id: string]: Room }>({})
  const videoRef = useRef<{ [room_id: string]: HTMLVideoElement }>({})
  //const roomIds = ["aaa", "bbb"]
  useEffect(() => {
    const ws = new WebSocket(SW_WSURL)
    const peer = new Peer({
      key: SKYWAY_APIKEY,
      debug: SKYWAY_DEBUG_LEVEL,
    })
    setWebSocket(ws)
    setSkywayPeer(peer)
    let sendFuncs: Array<Function> = []

    let isConnectWebSocket = false
    let isConnectSkywayPeer = false
    const nextRooms = {
      ...rooms,
    }
    roomIds.forEach((roomId) => {
      nextRooms[roomId] = { skyway_room: null, room_id: roomId }
    })
    setRooms(nextRooms)
    const sendFunc = () => {
      roomIds.forEach((roomId) => {
        ws.send(
          JSON.stringify({
            msg_type: "connect_receiver",
            room_id: roomId,
          })
        )
      })
    }
    if (ws.readyState == ws.OPEN) {
      sendFunc()
    } else {
      sendFuncs.push(sendFunc)
    }
    ws.addEventListener("open", (event) => {
      isConnectWebSocket = true
      if (!isConnectSkywayPeer) {
        return
      }
      for (const func of sendFuncs) {
        func()
      }
      sendFuncs = []
    })

    peer.on("open", () => {
      isConnectSkywayPeer = true
      if (!isConnectWebSocket) {
        return
      }
      for (const func of sendFuncs) {
        func()
      }
      sendFuncs = []
    })
    ws.addEventListener("message", (event) => {
      const message = JSON.parse(event.data)
      console.log(message)
      const peerId = message["peer_id"]
      const roomId = message["room_id"]
      const skywayRoomId = message["skyway_room_id"]

      console.log("joinroom")
      const skywayRoom = peer.joinRoom(skywayRoomId, {
        mode: "sfu",
      })
      const room: Room = { room_id: roomId, skyway_room: skywayRoom }
      const nextRooms = {
        ...rooms,
        [roomId]: room,
      }
      if (skywayRoom) {
        skywayRoom.on("stream", (stream: MediaStreamWithPeerId) => {
          const streamPeerId = stream.peerId
          console.log("on stream")
          if (streamPeerId == peerId) {
            const nextStreams = {
              ...castingStreams,
              [roomId]: stream,
            }
            setCastingStreams(nextStreams)
          }
        })
      }
      setRooms(nextRooms)
    })
    return () => {
      ws.close()
      peer.destroy()
    }
  }, [])

  useEffect(() => {
    roomIds.forEach((roomId) => {
      if (
        videoRef &&
        videoRef.current &&
        roomId in videoRef.current &&
        videoRef.current[roomId] &&
        roomId in castingStreams &&
        castingStreams[roomId]
      ) {
        if ("srcObject" in videoRef.current[roomId]) {
          if (videoRef.current[roomId].srcObject != castingStreams[roomId]) {
            videoRef.current[roomId].srcObject = castingStreams[roomId]
            try {
              videoRef.current[roomId].play()
            } catch (error) {}
            videoRef.current[roomId].volume = 0
          }
        } else {
          //videoRef.current.src = castingStreams["a"] // window.URL.createObjectURL(castingStream);
        }
      }
    })
  }, [castingStreams])
  return (
    <React.Fragment>
      {roomIds.map((roomId) => {
        return (
          <video
            width={400}
            ref={(el) => {
              if (el != null) {
                videoRef.current[roomId] = el
              }
            }}
            autoPlay
            playsInline
            key={roomId}
          ></video>
        )
      })}
    </React.Fragment>
  )
}

export default VideoCast
