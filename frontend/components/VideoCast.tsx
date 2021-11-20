import React, { FC, useEffect, useState, useRef, memo } from "react"
import Peer, { SfuRoom, MeshRoom } from "skyway-js"
import { SendMessage, WebRTCMessage } from "../types/webrtc-message"
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

const VideoCast: FC<Prop> = ({ roomIds }) => {
  const webSocket = useRef<WebSocket>()
  const [isWebSocketAvailable, setIsWebSocketAvailable] =
    useState<boolean>(false)
  const skywayPeer = useRef<Peer>()
  const [isPeerAvailable, setIsPeerAvailable] = useState<boolean>(false)

  useEffect(() => {
    webSocket.current = new WebSocket(SW_WSURL)
    skywayPeer.current = new Peer({
      key: SKYWAY_APIKEY,
      debug: SKYWAY_DEBUG_LEVEL,
    })
    webSocket.current.addEventListener("open", (event) => {
      if (webSocket.current) {
        setIsWebSocketAvailable(true)
      }
    })
    webSocket.current.addEventListener("close", (event) => {
      if (webSocket.current) {
        setIsWebSocketAvailable(false)
      }
    })

    skywayPeer.current.on("open", () => {
      console.log("opened skyway peer")
      setIsPeerAvailable(true)
    })
    return () => {
      webSocket.current?.close()
      skywayPeer.current?.destroy()
    }
  }, [])

  return (
    <React.Fragment>
      {roomIds.map((roomId) => {
        if (
          webSocket.current === undefined ||
          skywayPeer.current === undefined ||
          !isWebSocketAvailable ||
          !isPeerAvailable
        ) {
          return null
        }
        return (
          <RoomViewer
            roomId={roomId}
            ws={webSocket.current}
            peer={skywayPeer.current}
            key={roomId}
          />
        )
      })}
    </React.Fragment>
  )
}

type RoomViewerProps = {
  roomId: string
  ws: WebSocket
  peer: Peer
}

const RoomViewer: FC<RoomViewerProps> = ({ roomId, ws, peer }) => {
  const [skywayRoom, setSkywayRoom] = useState<SfuRoom | MeshRoom | null>(null)
  const [peerId, setPeerId] = useState<string | null>(null)

  useEffect(() => {
    const onMessage = (event: WebSocketEventMap["message"]) => {
      const message: SendMessage = JSON.parse(event.data)
      if (message.room_id !== roomId) {
        return
      }
      console.log("ruu message", message)
      const peerId = message.peer_id
      const skywayRoomId = message.skyway_room_id

      console.log("joinroom")
      const skywayRoom = peer.joinRoom(skywayRoomId, {
        mode: "sfu",
      })
      setSkywayRoom(skywayRoom)
      setPeerId(peerId)
    }
    ws.addEventListener("message", onMessage)

    const connectMessage: WebRTCMessage = {
      msg_type: "connect_receiver",
      room_id: roomId,
    }
    ws.send(JSON.stringify(connectMessage))
    console.log("ruu connect_receiver", roomId)

    return () => {
      ws.removeEventListener("message", onMessage)
      const exitMessage = {
        msg_type: "exit_room",
        room_id: roomId,
      }
      ws.send(JSON.stringify(exitMessage))
      console.log("ruu close send", roomId)
    }
  }, [])
  if (peerId === null) {
    return <p>PREPARING VIDEO</p>
  }
  return <SkywayRoomViewer room={skywayRoom} peerId={peerId} />
}

type SkywayRoomViewerProps = {
  room: SfuRoom | MeshRoom
  peerId: string
}

const SkywayRoomViewer: FC<SkywayRoomViewerProps> = ({ room, peerId }) => {
  const ref = useRef<HTMLVideoElement>(null)
  const [castingStream, setCastingStream] = useState<MediaStreamWithPeerId>()
  useEffect(() => {
    room.on("stream", (stream: MediaStreamWithPeerId) => {
      const streamPeerId = stream.peerId
      console.log("ruu on stream", stream)
      if (streamPeerId !== peerId) {
        return
      }
      setCastingStream(stream)
    })
  }, [room, setCastingStream])
  useEffect(() => {
    const video = ref.current
    if (video === null || castingStream === undefined) {
      return
    }
    if (video.srcObject !== castingStream) {
      video.srcObject = castingStream
      try {
        video.play()
      } catch (err) {
        console.error(err)
      }
      video.volume = 0
    }
  }, [castingStream])
  return <video width={400} ref={ref} autoPlay playsInline></video>
}

const MemoizedSkywayRoomViewer = memo(SkywayRoomViewer)

export default VideoCast
