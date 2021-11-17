import React, { FC } from "react"
import Rail from "./svgParts/Rail"
import Station from "./svgParts/Station"
import SwitchPoint from "./svgParts/SwitchPoint"
import StopPoint from "./svgParts/StopPoint"
import { TrainData } from "../types/svgPartsTypes"
import { StopRailId } from "../types/websocket-messages"

interface Prop {
  datas: {
    stop: Record<StopRailId, boolean>
    stop1: boolean
    stop2: boolean
    stop3: boolean
    stop4: boolean
    switch1: boolean
    switch2: boolean
    train1: TrainData
  }
}

const RailroadMap: FC<Prop> = ({
  datas: { stop, stop1, stop2, stop3, stop4, switch1, switch2, train1 },
}) => {
  return (
    <svg width={640} height={480} viewBox="0 0 640 480">
      <rect x={0} y={0} width={640} height={480} fill="lightgray" />

      <Station name="東京" position={{ x: 120, y: 120 }} />
      <Station name="札幌" position={{ x: 520, y: 120 }} />
      <Station name="那覇" position={{ x: 320, y: 360 }} />

      <SwitchPoint
        position={{ x: 320 + 20, y: 120 - 20 }}
        fromAngle={180}
        leftOutAngle={0}
        rightOutAngle={90}
        isLeft={switch1}
      />
      <SwitchPoint
        position={{ x: 320 - 20, y: 120 + 20 }}
        fromAngle={180}
        leftOutAngle={0}
        rightOutAngle={90}
        isLeft={switch2}
      />

      <Rail
        startPosition={{ x: 120 - 50, y: 120 - 20 }}
        endPosition={{ x: 320 + 10, y: 120 - 20 }}
        trains={[]}
      />
      <Rail
        startPosition={{ x: 120 - 50, y: 120 + 20 }}
        endPosition={{ x: 320 - 30, y: 120 + 20 }}
        trains={[train1]}
      />
      <Rail
        startPosition={{ x: 320 + 30, y: 120 - 20 }}
        endPosition={{ x: 520 + 50, y: 120 - 20 }}
        trains={[]}
      />
      <Rail
        startPosition={{ x: 320 - 10, y: 120 + 20 }}
        endPosition={{ x: 520 + 50, y: 120 + 20 }}
        trains={[]}
      />
      <Rail
        startPosition={{ x: 320 - 20, y: 120 + 30 }}
        endPosition={{ x: 320 - 20, y: 360 - 10 }}
        trains={[]}
      />
      <Rail
        startPosition={{ x: 320 + 20, y: 120 - 10 }}
        endPosition={{ x: 320 + 20, y: 360 - 10 }}
        trains={[]}
      />

      <StopPoint position={{ x: 120, y: 120 - 20 }} isStop={stop1} />
      <StopPoint position={{ x: 120, y: 120 + 20 }} isStop={stop2} />
      <StopPoint position={{ x: 520, y: 120 - 20 }} isStop={stop3} />
      <StopPoint position={{ x: 520, y: 120 + 20 }} isStop={stop4} />

      <StopPoint
        position={{ x: 10 + 20 * 0, y: 10 + 20 * 0 }}
        isStop={stop.motoyawata_s1}
      />
      <StopPoint
        position={{ x: 10 + 20 * 0, y: 10 + 20 * 1 }}
        isStop={stop.motoyawata_s2}
      />
      <StopPoint
        position={{ x: 10 + 20 * 0, y: 10 + 20 * 2 }}
        isStop={stop.motoyawata_s3}
      />
      <StopPoint
        position={{ x: 10 + 20 * 0, y: 10 + 20 * 3 }}
        isStop={stop.motoyawata_s4}
      />
      <StopPoint
        position={{ x: 10 + 20 * 0, y: 10 + 20 * 4 }}
        isStop={stop.motoyawata_s5}
      />
      <StopPoint
        position={{ x: 10 + 20 * 0, y: 10 + 20 * 5 }}
        isStop={stop.motoyawata_s6}
      />
      <StopPoint
        position={{ x: 10 + 20 * 1, y: 10 + 20 * 0 }}
        isStop={stop.iwamotocho_s1}
      />
      <StopPoint
        position={{ x: 10 + 20 * 1, y: 10 + 20 * 1 }}
        isStop={stop.iwamotocho_s2}
      />
      <StopPoint
        position={{ x: 10 + 20 * 1, y: 10 + 20 * 2 }}
        isStop={stop.iwamotocho_s4}
      />
      <StopPoint
        position={{ x: 10 + 20 * 2, y: 10 + 20 * 0 }}
        isStop={stop.kudanshita_s5}
      />
      <StopPoint
        position={{ x: 10 + 20 * 2, y: 10 + 20 * 1 }}
        isStop={stop.kudanshita_s6}
      />
      <StopPoint
        position={{ x: 10 + 20 * 3, y: 10 + 20 * 0 }}
        isStop={stop.sasazuka_s1}
      />
      <StopPoint
        position={{ x: 10 + 20 * 3, y: 10 + 20 * 1 }}
        isStop={stop.sasazuka_s2}
      />
      <StopPoint
        position={{ x: 10 + 20 * 3, y: 10 + 20 * 2 }}
        isStop={stop.sasazuka_s3}
      />
      <StopPoint
        position={{ x: 10 + 20 * 3, y: 10 + 20 * 3 }}
        isStop={stop.sasazuka_s3}
      />
      <StopPoint
        position={{ x: 10 + 20 * 3, y: 10 + 20 * 4 }}
        isStop={stop.sasazuka_s5}
      />
      <StopPoint
        position={{ x: 10 + 20 * 4, y: 10 + 20 * 0 }}
        isStop={stop.meidaimae_s1}
      />
      <StopPoint
        position={{ x: 10 + 20 * 4, y: 10 + 20 * 1 }}
        isStop={stop.meidaimae_s2}
      />
      <StopPoint
        position={{ x: 10 + 20 * 5, y: 10 + 20 * 0 }}
        isStop={stop.chofu_s1}
      />
      <StopPoint
        position={{ x: 10 + 20 * 5, y: 10 + 20 * 1 }}
        isStop={stop.chofu_s2}
      />
      <StopPoint
        position={{ x: 10 + 20 * 5, y: 10 + 20 * 2 }}
        isStop={stop.chofu_s3}
      />
      <StopPoint
        position={{ x: 10 + 20 * 5, y: 10 + 20 * 3 }}
        isStop={stop.chofu_s4}
      />
      <StopPoint
        position={{ x: 10 + 20 * 5, y: 10 + 20 * 4 }}
        isStop={stop.chofu_s5}
      />
      <StopPoint
        position={{ x: 10 + 20 * 5, y: 10 + 20 * 5 }}
        isStop={stop.chofu_s6}
      />
      <StopPoint
        position={{ x: 10 + 20 * 6, y: 10 + 20 * 0 }}
        isStop={stop.kitano_s1}
      />
      <StopPoint
        position={{ x: 10 + 20 * 6, y: 10 + 20 * 1 }}
        isStop={stop.kitano_s2}
      />
      <StopPoint
        position={{ x: 10 + 20 * 6, y: 10 + 20 * 2 }}
        isStop={stop.kitano_s3}
      />
      <StopPoint
        position={{ x: 10 + 20 * 6, y: 10 + 20 * 3 }}
        isStop={stop.kitano_s4}
      />
      <StopPoint
        position={{ x: 10 + 20 * 6, y: 10 + 20 * 4 }}
        isStop={stop.kitano_s5}
      />
      <StopPoint
        position={{ x: 10 + 20 * 6, y: 10 + 20 * 5 }}
        isStop={stop.kitano_s6}
      />
      <StopPoint
        position={{ x: 10 + 20 * 7, y: 10 + 20 * 0 }}
        isStop={stop.takao_s1}
      />
      <StopPoint
        position={{ x: 10 + 20 * 7, y: 10 + 20 * 1 }}
        isStop={stop.takao_s2}
      />
    </svg>
  )
}

export default RailroadMap
