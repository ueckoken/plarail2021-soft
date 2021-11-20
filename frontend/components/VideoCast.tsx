import React, { FC, useEffect, useState, useRef } from "react"
import Peer, { SfuRoom, MeshRoom } from "skyway-js"
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

interface Room {
  skyway_room: SfuRoom | MeshRoom | null
  //video_ref: HTMLAnchorElement | null;
  room_id: string | null
}
const VideoCast: FC<Prop> = ({ roomIds }) => {
  const [castingStreams, setCastingStreams] = useState<{
    [room_id: string]: MediaStream | null
  }>({})
  const webSocket = useRef<WebSocket>()
  const skywayPeer = useRef<Peer>()
  const [rooms, setRooms] = useState<{ [room_id: string]: Room }>({})
  const videoRef = useRef<HTMLVideoElement[]>([])
  const sendFuncs = useRef<Function[]>([])
  useEffect(() => {
    webSocket.current = new WebSocket(SW_WSURL)
    skywayPeer.current = new Peer({
      key: SKYWAY_APIKEY,
      debug: SKYWAY_DEBUG_LEVEL,
    })
    let isConnectWebSocket = false
    let isConnectSkywayPeer = false
    webSocket.current.addEventListener("open", (event) => {
      isConnectWebSocket = true
      if (!isConnectSkywayPeer) {
        return
      }
      for (const func of sendFuncs.current) {
        func()
      }
      sendFuncs.current = []
    })

    skywayPeer.current.on("open", () => {
      isConnectSkywayPeer = true
      if (!isConnectWebSocket) {
        return
      }
      for (const func of sendFuncs.current) {
        func()
      }
      sendFuncs.current = []
    })
    webSocket.current.addEventListener("message", (event) => {
      const message = JSON.parse(event.data)
      console.log("ruu message", message)
      const peerId = message["peer_id"]
      const roomId = message["room_id"]
      const skywayRoomId = message["skyway_room_id"]

      console.log("joinroom")
      if (skywayPeer.current == undefined) {
        return
      }
      const skywayRoom = skywayPeer.current.joinRoom(skywayRoomId, {
        mode: "sfu",
      })
      const room: Room = { room_id: roomId, skyway_room: skywayRoom }
      let nextRooms: { [room_id: string]: Room } = {}
      Object.keys(rooms).forEach((id) => {
        nextRooms[id] = rooms[id]
      })
      nextRooms[roomId] = room
      /*const nextRooms = {
        ...rooms,
        [roomId]: room,
      }*/
      if (skywayRoom) {
        skywayRoom.on("stream", (stream: MediaStreamWithPeerId) => {
          const streamPeerId = stream.peerId
          console.log("ruu on stream")
          if (streamPeerId == peerId) {
            const nextStreams = {
              ...castingStreams,
              [roomId]: stream,
            }
            console.log(
              "ruu",
              Object.keys(castingStreams),
              Object.keys(nextStreams)
            )
            setCastingStreams(nextStreams)
          }
        })
      }
      console.log("ruu", rooms, nextRooms)
      setRooms(nextRooms)
    })
    return () => {
      if (webSocket.current) {
        webSocket.current.close()
      }
      if (skywayPeer.current) {
        skywayPeer.current.destroy()
      }
    }
  }, [])

  useEffect(() => {
    if (JSON.stringify(Object.keys(rooms)) == JSON.stringify(roomIds)) {
      return
    }
    console.log("ruu", roomIds)
    const nextRooms: { [room_id: string]: Room } = {}
    console.log("ruu", Object.keys(rooms))
    Object.keys(rooms).forEach((roomId) => {
      if (roomIds.includes(roomId)) {
        nextRooms[roomId] = rooms[roomId]
      } else {
        if (rooms[roomId]["skyway_room"]) {
          rooms[roomId]["skyway_room"]!.close()
          console.log("ruu close", roomId)
        }

        const sendfunc = () => {
          if (!webSocket.current) {
            return
          }
          webSocket.current.send(
            JSON.stringify({
              msg_type: "exit_room",
              room_id: roomId,
            })
          )
          console.log("ruu close send", roomId)
        }
        if (
          webSocket.current &&
          webSocket.current.readyState == webSocket.current.OPEN &&
          skywayPeer.current
        ) {
          sendfunc()
        } else {
          sendFuncs.current.push(sendfunc)
        }
      }
    })
    roomIds.forEach((roomId) => {
      if (!Object.keys(rooms).includes(roomId)) {
        const sendfunc = () => {
          if (!webSocket.current) {
            return
          }
          webSocket.current.send(
            JSON.stringify({
              msg_type: "connect_receiver",
              room_id: roomId,
            })
          )
          console.log("ruu connect_receiver", roomId)
        }
        if (
          webSocket.current &&
          webSocket.current.readyState == webSocket.current.OPEN &&
          skywayPeer.current
        ) {
          sendfunc()
        } else {
          sendFuncs.current.push(sendfunc)
        }
      }
      nextRooms[roomId] = { room_id: roomId, skyway_room: null }
    })
    console.log("ruu", rooms, nextRooms)
    setRooms(nextRooms)
  }, [roomIds])

  useEffect(() => {
    roomIds.forEach((roomId, index) => {
      if (videoRef && videoRef.current && videoRef.current[index]) {
        if ("srcObject" in videoRef.current[index]) {
          if (videoRef.current[index].srcObject != castingStreams[roomId]) {
            videoRef.current[index].srcObject = castingStreams[roomId]
            try {
              videoRef.current[index].play()
            } catch (error) {}
            videoRef.current[index].volume = 0
          }
        } else {
          //videoRef.current.src = castingStreams["a"] // window.URL.createObjectURL(castingStream);
        }
      }
    })
  })
  return (
    <React.Fragment>
      {roomIds.map((roomId, index) => {
        return (
          <video
            width={400}
            ref={(el) => {
              if (el != null) {
                videoRef.current[index] = el
              }
            }}
            autoPlay
            playsInline
            key={index}
          ></video>
        )
      })}
    </React.Fragment>
  )
}

export default VideoCast
