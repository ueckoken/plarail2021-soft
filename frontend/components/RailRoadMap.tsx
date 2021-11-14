import React, { FC } from "react"
import Rail from "./svgParts/Rail"
import Station from "./svgParts/Station"
import SwitchPoint from "./svgParts/SwitchPoint"
import StopPoint from "./svgParts/StopPoint"
import { TrainData } from "../types/svgPartsTypes"

interface Prop {
  datas: {
    stop1: boolean
    switch1: boolean
    train1: TrainData
    station1: boolean
    station2: boolean
    station3: boolean
  }
}

const RailroadMap: FC<Prop> = ({
  datas: { stop1, switch1, train1, station1, station2, station3 },
}) => {
  return (
    <svg width={640} height={480} viewBox="0 0 640 480">
      <rect x={0} y={0} width={640} height={480} fill="lightgray" />

      <Station name="東京" position={{ x: 120, y: 120 }} isStop={station1} />
      <Station name="札幌" position={{ x: 520, y: 120 }} isStop={station2} />
                        <Station name="那覇" position={{ x: 320, y: 360 }} isStop={station3} />

      <SwitchPoint        position={{ x: 320, y: 120 }}        fromAngle={180}
        leftOutAngle={0}
        rightOutAngle={90}
        isLeft={switch1}
      />

      <Rail
        startPosition={{ x: 120 + 50, y: 120 }}
        endPosition={{ x: 320 - 10, y: 120 }}
        trains={[train1]}
      />
      <Rail
        startPosition={{ x: 320 + 10, y: 120 }}
        endPosition={{ x: 520 - 50, y: 120 }}
        trains={[]}
      />
      <Rail
        startPosition={{ x: 320, y: 120 + 10 }}
        endPosition={{ x: 320, y: 360 - 10 }}
        trains={[]}
      />

      <StopPoint position={{ x: 320, y: 240 }} isStop={stop1} />
    </svg>
  )
}

export default RailroadMap
