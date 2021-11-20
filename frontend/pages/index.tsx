import type { NextPage } from "next"
import Head from "next/head"
import styles from "../styles/Home.module.css"
import RailroadMap from "../components/RailRoadMap"
import VideoCast from "../components/VideoCast"
import { useEffect, useRef, useState } from "react"
import {
  bunkiRailId,
  BunkiRailId,
  Message,
  stopRailId,
  StopRailId,
} from "../types/control-messages"
import SpeedMeter from "../components/svgParts/SpeedMeter"
import { SpeedMessage, TrainId } from "../types/speed-messages"
import ReverseHandle from "../components/svgParts/ReverseHandle"

// OFF: false, ON: trueと対応
type StopPointState = Record<StopRailId, boolean>
const INITIAL_STOP_POINT_STATE: StopPointState = {
  motoyawata_s1: false,
  motoyawata_s2: false,
  iwamotocho_s1: false,
  iwamotocho_s2: false,
  iwamotocho_s4: false,
  kudanshita_s5: false,
  kudanshita_s6: false,
  sasazuka_s1: false,
  sasazuka_s2: false,
  sasazuka_s3: false,
  sasazuka_s4: false,
  sasazuka_s5: false,
  meidaimae_s1: false,
  meidaimae_s2: false,
  chofu_s1: false,
  chofu_s2: false,
  chofu_s3: false,
  chofu_s4: false,
  chofu_s5: false,
  chofu_s6: false,
  kitano_s1: false,
  kitano_s2: false,
  kitano_s3: false,
  kitano_s4: false,
  kitano_s5: false,
  kitano_s6: false,
  kitano_s7: false,
  takao_s1: false,
  takao_s2: false,
}

type SwitchPointState = Record<BunkiRailId, boolean>
const INITIAL_SWITCH_POINT_STATE: SwitchPointState = {
  iwamotocho_b1: false,
  iwamotocho_b2: false,
  iwamotocho_b3: false,
  iwamotocho_b4: false,
  sasazuka_b1: false,
  sasazuka_b2: false,
  chofu_b1: false,
  chofu_b2: false,
  chofu_b3: false,
  chofu_b4: false,
  chofu_b5: false,
  kitano_b1: false,
  kitano_b2: false,
  kitano_b3: false,
}

type SpeedState = Record<TrainId, number>
const INITIAL_SPEED_STATE: SpeedState = {
  TAKAO: 0,
  CHICHIBU: 0,
  HAKONE: 0,
  OKUTAMA: 0,
  NIKKO: 0,
  ENOSHIMA: 0,
  KAMAKURA: 0,
  YOKOSUKA: 0,
}
const INITIAL_SELECTED_TRAIN_ID: TrainId = "TAKAO"

const Home: NextPage = () => {
  const stationWs = useRef<WebSocket>()
  const [stopPointState, setStopPointState] = useState<StopPointState>(
    INITIAL_STOP_POINT_STATE
  )
  const [switchPointState, setSwitchPointState] = useState<SwitchPointState>(
    INITIAL_SWITCH_POINT_STATE
  )
  const [isLeftSwichPoint1, setIsLeftSwitchPoint1] = useState<boolean>(true)
  const [isLeftSwichPoint2, setIsLeftSwitchPoint2] = useState<boolean>(true)
  const [trainPosition1, setTrainPosition1] = useState<number>(0.4)

  const speedWs = useRef<WebSocket>()
  const [speedState, setSpeedState] = useState<SpeedState>(INITIAL_SPEED_STATE)
  const [selectedTrainId, setSelectedTrainId] = useState<TrainId>(
    INITIAL_SELECTED_TRAIN_ID
  )
  const [isBack, setIsBack] = useState<boolean>(false)

  const [roomIds, setRoomIds] = useState<string[]>(["aaa"])

  useEffect(() => {
    const ws = new WebSocket("wss://speed.chofufes2021.gotti.dev/speed")
    speedWs.current = ws
    ws.addEventListener("open", (e) => {
      console.log("opened")
    })
    ws.addEventListener("message", (e) => {
      const message: SpeedMessage = JSON.parse(e.data)
      console.log(message)
      setSpeedState((previousState) => ({
        ...previousState,
        [message.train_name]: message.speed,
      }))
    })
    ws.addEventListener("error", (e) => {
      console.log("error occured")
      console.log(e)
    })
    ws.addEventListener("close", (e) => {
      console.log("closed")
      console.log(e)
    })
    return () => {
      ws.close()
    }
  }, [])

  useEffect(() => {
    const ws = new WebSocket("wss://control.chofufes2021.gotti.dev/ws")
    stationWs.current = ws
    ws.addEventListener("open", (e) => {
      console.log("opened")
      console.log(e)
    })
    ws.addEventListener("message", (e) => {
      console.log("recieved message")
      console.log(e)
      const message: Message = JSON.parse(e.data)
      console.log(message)
      if (message.station_name === "unknown" || message.state === "UNKNOWN") {
        return
      }
      if (stopRailId.is(message.station_name)) {
        setStopPointState((previousStopPointState) => ({
          ...previousStopPointState,
          [message.station_name]: message.state === "ON",
        }))
        return
      }
      if (bunkiRailId.is(message.station_name)) {
        setSwitchPointState((previousSwitchPointState) => ({
          ...previousSwitchPointState,
          [message.station_name]: message.state === "ON",
        }))
      }
    })
    ws.addEventListener("error", (e) => {
      console.log("error occured")
      console.log(e)
    })
    ws.addEventListener("close", (e) => {
      console.log("closed")
      console.log(e)
    })
    return () => {
      ws.close()
    }
  }, [])

  useEffect(() => {
    const intervalId = setInterval(() => {
      const tmpTrainPosition1 = trainPosition1 + 0.01
      setTrainPosition1(tmpTrainPosition1 <= 1 ? tmpTrainPosition1 : 0)
    }, 20)
    return () => clearInterval(intervalId)
  })

  return (
    <div className={styles.container}>
      <Head>
        <title>工研&times;鉄研プラレール展示 操作ページ</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/kokenLogo.ico" />
      </Head>

      <header>
        <h1 className={styles.title}>
          工研&times;鉄研プラレール展示 操作ページ
        </h1>
      </header>

      <main className={styles.main}>
        <section>
          <h2>映像部分</h2>
          <div>(video予定地)</div>
          <VideoCast roomIds={roomIds} />
          <button
            onClick={() => {
              setRoomIds(["aaa", "bbb"])
            }}
          >
            aaa
          </button>
          <button
            onClick={() => {
              setRoomIds(["bbb", "aaa"])
            }}
          >
            bbb
          </button>
          <button
            onClick={() => {
              setRoomIds(["ccc", "bbb"])
            }}
          >
            ccc
          </button>
        </section>

        <section>
          <h2>地図部分</h2>
          <RailroadMap
            datas={{
              stop: stopPointState,
              switchState: switchPointState,
              train1: {
                positionScale: trainPosition1,
                id: "koken",
              },
            }}
          />
        </section>

        <section>
          <h2>操作部分</h2>
          <svg width="100%" viewBox="0 0 200 100">
            <rect x={0} y={0} width={200} height={100} fill="dimgrey" />
            <SpeedMeter
              cx={80}
              cy={40}
              r={30}
              max={100}
              value={Math.abs(speedState[selectedTrainId])}
            />
            <ReverseHandle
              cx={150}
              cy={50}
              r={3}
              isBack={isBack}
              onChange={(nextIsBack) => {
                setIsBack(nextIsBack)
              }}
            />
          </svg>
          <input
            type="range"
            min={0}
            defaultValue={0}
            max={100}
            step={5}
            onChange={(e) => {
              const a = isBack ? -1 : 1 // 前進なら1、逆進なら-1になる係数
              const message: SpeedMessage = {
                train_name: selectedTrainId,
                speed: a * Number.parseInt(e.target.value),
              }
              speedWs.current?.send(JSON.stringify(message))
            }}
          />
          <button
            onClick={() => {
              const message: Message = {
                station_name: "chofu_b1",
                state: switchPointState.chofu_b1 ? "OFF" : "ON",
              }
              stationWs.current?.send(JSON.stringify(message))
            }}
          >
            chofu_b1切り替え
          </button>
          <button
            onClick={() => {
              const message: Message = {
                station_name: "chofu_b2",
                state: switchPointState.chofu_b2 ? "OFF" : "ON",
              }
              stationWs.current?.send(JSON.stringify(message))
            }}
          >
            chofu_b2切り替え
          </button>
          <button
            onClick={() => {
              const message: Message = {
                station_name: "chofu_b3",
                state: switchPointState.chofu_b3 ? "OFF" : "ON",
              }
              stationWs.current?.send(JSON.stringify(message))
            }}
          >
            chofu_b3切り替え
          </button>
          <button
            onClick={() => {
              const message: Message = {
                station_name: "chofu_b4",
                state: switchPointState.chofu_b4 ? "OFF" : "ON",
              }
              stationWs.current?.send(JSON.stringify(message))
            }}
          >
            chofu_b4切り替え
          </button>
          <button
            onClick={() => {
              const message: Message = {
                station_name: "chofu_b5",
                state: switchPointState.chofu_b5 ? "OFF" : "ON",
              }
              stationWs.current?.send(JSON.stringify(message))
            }}
          >
            chofu_b5切り替え
          </button>
          <button
            onClick={() => {
              const message: Message = {
                station_name: "chofu_s1",
                state: stopPointState.chofu_s1 ? "OFF" : "ON",
              }
              stationWs.current?.send(JSON.stringify(message))
            }}
          >
            chofu_s1切り替え
          </button>
          <button
            onClick={() => {
              const message: Message = {
                station_name: "chofu_s2",
                state: stopPointState.chofu_s2 ? "OFF" : "ON",
              }
              stationWs.current?.send(JSON.stringify(message))
            }}
          >
            chofu_s2切り替え
          </button>
          <button
            onClick={() => {
              const message: Message = {
                station_name: "chofu_s3",
                state: stopPointState.chofu_s3 ? "OFF" : "ON",
              }
              stationWs.current?.send(JSON.stringify(message))
            }}
          >
            chofu_s3切り替え
          </button>
          <button
            onClick={() => {
              const message: Message = {
                station_name: "chofu_s4",
                state: stopPointState.chofu_s4 ? "OFF" : "ON",
              }
              stationWs.current?.send(JSON.stringify(message))
            }}
          >
            chofu_s4切り替え
          </button>
        </section>
      </main>

      <footer className={styles.footer}>
        <a
          href="https://www.koken.club.uec.ac.jp/"
          target="_blank"
          rel="noopener noreferrer"
        >
          &copy;2021 電気通信大学工学研究部
        </a>
      </footer>
    </div>
  )
}

export default Home
