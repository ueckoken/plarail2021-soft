import React, { FC } from "react"
import Rail from "./svgParts/Rail"
import Platform from "./svgParts/Platform"
import SwitchPoint from "./svgParts/SwitchPoint"
import StopPoint from "./svgParts/StopPoint"
import { TrainData } from "../types/svgPartsTypes"
import { BunkiRailId, StopRailId } from "../types/websocket-messages"

interface Prop {
  datas: {
    stop: Record<StopRailId, boolean>
    stop1: boolean
    stop2: boolean
    stop3: boolean
    stop4: boolean
    stop5: boolean
    stop6: boolean
    stop7: boolean
    stop8: boolean
    stop9: boolean
    stop10: boolean
    stop11: boolean
    stop12: boolean
    stop13: boolean
    stop14: boolean
    stop15: boolean
    stop16: boolean
    stop17: boolean
    stop18: boolean
    stop19: boolean
    stop20: boolean
    stop21: boolean
    stop22: boolean
    stop23: boolean
    stop24: boolean
    stop25: boolean
    stop26: boolean
    switchState: Record<BunkiRailId, boolean>
    switch1: boolean
    switch2: boolean
    switch3: boolean
    switch4: boolean
    switch5: boolean
    switch6: boolean
    switch7: boolean
    switch8: boolean
    switch9: boolean
    switch10: boolean
    switch11: boolean
    train1: TrainData
  }
}

const RailroadMap: FC<Prop> = ({
  datas: {
    stop,
    stop1,
    stop2,
    stop3,
    stop4,
    stop5,
    stop6,
    stop7,
    stop8,
    stop9,
    stop10,
    stop11,
    stop12,
    stop13,
    stop14,
    stop15,
    stop16,
    stop17,
    stop18,
    stop19,
    stop20,
    stop21,
    stop22,
    stop23,
    stop24,
    stop25,
    stop26,
    switchState,
    switch1,
    switch2,
    switch3,
    switch4,
    switch5,
    switch6,
    switch7,
    switch8,
    switch9,
    switch10,
    switch11,
    train1,
  },
}) => {
  return (
    <svg width={1120} height={620} viewBox="0 0 1120 620">
      <rect x={0} y={0} width={1120} height={620} fill="lightgray" />

      <Platform name="京王八王子" position={{ x: 60, y: 70 }} />
      <Platform name="北野1" position={{ x: 240, y: 40 }} />
      <Platform name="北野2" position={{ x: 240, y: 100 }} />
      <Platform name="調布1" position={{ x: 520, y: 40 }} />
      <Platform name="調布2" position={{ x: 520, y: 140 }} />
      <Platform name="明大前1" position={{ x: 700, y: 40 }} />
      <Platform name="明大前2" position={{ x: 700, y: 100 }} />
      <Platform name="笹塚1" position={{ x: 880, y: 40 }} />
      <Platform name="笹塚2" position={{ x: 880, y: 100 }} />
      <Platform name="新宿" position={{ x: 1060, y: 70 }} />

      <Platform name="高尾" position={{ x: 150, y: 200 }} isHorizontal={false} />
      <Platform name="高尾山口" position={{ x: 150, y: 380 }} isHorizontal={false} />
      <Platform name="若葉台1" position={{ x: 320, y: 260 }} isHorizontal={false} />
      <Platform name="若葉台2" position={{ x: 380, y: 260 }} isHorizontal={false} />
      <Platform name="橋本" position={{ x: 350, y: 380 }} isHorizontal={false} />

      <Platform name="九段下" position={{ x: 970, y: 180 }} isHorizontal={false} />
      <Platform name="岩本町1" position={{ x: 910, y: 380 }} isHorizontal={false} />
      <Platform name="岩本町2" position={{ x: 970, y: 380 }} isHorizontal={false} />
      <Platform name="本八幡" position={{ x: 970, y: 560 }} isHorizontal={false} />

      <Rail
        positions={[
          { x: 120, y: 120 },
          { x: 20, y: 120 },
          { x: 20, y: 20 },
          { x: 140, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 120, y: 120 },
          { x: 20, y: 120 },
          { x: 20, y: 20 },
          { x: 140, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 140, y: 20 },
          { x: 160, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 140, y: 20 },
          { x: 140, y: 100 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 140, y: 100 },
          { x: 120, y: 120 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 120, y: 120 },
          { x: 120, y: 240 },
          { x: 150, y: 270 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 150, y: 270 },
          { x: 150, y: 310 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 150, y: 310 },
          { x: 120, y: 340 },
          { x: 120, y: 420 },
          { x: 180, y: 420 },
          { x: 180, y: 340 },
          { x: 150, y: 310 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 150, y: 270 },
          { x: 180, y: 240 },
          { x: 180, y: 120 },
          { x: 320, y: 120 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 160, y: 20 },
          { x: 320, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 160, y: 20 },
          { x: 200, y: 60 },
          { x: 280, y: 60 },
          { x: 320, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 320, y: 20 },
          { x: 380, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 140, y: 100 },
          { x: 160, y: 80 },
          { x: 280, y: 80 },
          { x: 320, y: 120 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 320, y: 120 },
          { x: 380, y: 120 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 380, y: 120 },
          { x: 420, y: 160 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 420, y: 160 },
          { x: 440, y: 160 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 440, y: 160 },
          { x: 480, y: 120 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 380, y: 120 },
          { x: 480, y: 120 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 480, y: 120 },
          { x: 600, y: 120 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 440, y: 160 },
          { x: 560, y: 160 },
          { x: 600, y: 120 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 380, y: 20 },
          { x: 480, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 380, y: 20 },
          { x: 420, y: 60 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 420, y: 60 },
          { x: 440, y: 60 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 440, y: 60 },
          { x: 480, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 480, y: 20 },
          { x: 600, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 440, y: 60 },
          { x: 560, y: 60 },
          { x: 600, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 420, y: 60 },
          { x: 320, y: 160 },
          { x: 320, y: 200 },
          { x: 340, y: 220 },
          { x: 340, y: 300 },
          { x: 320, y: 320 },
          { x: 320, y: 420 },
          { x: 380, y: 420 },
          { x: 380, y: 320 },
          { x: 360, y: 300 },
          { x: 360, y: 220 },
          { x: 380, y: 200 },
          { x: 380, y: 160 },
          { x: 420, y: 160 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 600, y: 20 },
          { x: 620, y: 20 },
          { x: 660, y: 60 },
          { x: 740, y: 60 },
          { x: 780, y: 20 },
          { x: 800, y: 20 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 600, y: 120 },
          { x: 620, y: 120 },
          { x: 660, y: 80 },
          { x: 740, y: 80 },
          { x: 780, y: 120 },
          { x: 800, y: 120 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 800, y: 20 },
          { x: 840, y: 60 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 800, y: 120 },
          { x: 840, y: 80 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 800, y: 20 },
          { x: 1100, y: 20 },
          { x: 1100, y: 120 },
          { x: 800, y: 120 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 840, y: 60 },
          { x: 800, y: 60 },
          { x: 800, y: 80 },
          { x: 840, y: 80 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 840, y: 60 },
          { x: 1000, y: 60 },
          { x: 1000, y: 280 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 840, y: 80 },
          { x: 940, y: 80 },
          { x: 940, y: 280 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 940, y: 280 },
          { x: 880, y: 340 },
          { x: 880, y: 420 },
          { x: 940, y: 480 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 940, y: 280 },
          { x: 940, y: 340 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 1000, y: 280 },
          { x: 940, y: 340 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 940, y: 340 },
          { x: 940, y: 420 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 940, y: 420 },
          { x: 1000, y: 480 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 1000, y: 280 },
          { x: 1000, y: 480 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 940, y: 420 },
          { x: 940, y: 480 }
        ]}
        trains={[]}
      />
      <Rail
        positions={[
          { x: 940, y: 480 },
          { x: 940, y: 600 },
          { x: 1000, y: 600 },
          { x: 1000, y: 480 }
        ]}
        trains={[]}
      />

      <StopPoint position={{ x: 120, y: 20 }} isStop={stop1} />
      <StopPoint position={{ x: 240, y: 20 }} isStop={stop1} />
      <StopPoint position={{ x: 240, y: 60 }} isStop={stop1} />
      <StopPoint position={{ x: 240, y: 80 }} isStop={stop1} />
      <StopPoint position={{ x: 240, y: 120 }} isStop={stop1} />
      <StopPoint position={{ x: 520, y: 20 }} isStop={stop1} />
      <StopPoint position={{ x: 520, y: 60 }} isStop={stop1} />
      <StopPoint position={{ x: 520, y: 120 }} isStop={stop1} />
      <StopPoint position={{ x: 520, y: 160 }} isStop={stop1} />
      <StopPoint position={{ x: 700, y: 60 }} isStop={stop1} />
      <StopPoint position={{ x: 700, y: 80 }} isStop={stop1} />
      <StopPoint position={{ x: 800, y: 70 }} isStop={stop1} />
      <StopPoint position={{ x: 880, y: 20 }} isStop={stop1} />
      <StopPoint position={{ x: 880, y: 60 }} isStop={stop1} />
      <StopPoint position={{ x: 880, y: 80 }} isStop={stop1} />
      <StopPoint position={{ x: 880, y: 120 }} isStop={stop1} />

      <StopPoint position={{ x: 120, y: 140 }} isStop={stop1} />
      <StopPoint position={{ x: 180, y: 200 }} isStop={stop1} />
      <StopPoint position={{ x: 120, y: 380 }} isStop={stop1} />

      <StopPoint position={{ x: 320, y: 180 }} isStop={stop1} />

      <StopPoint position={{ x: 940, y: 180 }} isStop={stop1} />
      <StopPoint position={{ x: 1000, y: 180 }} isStop={stop1} />
      <StopPoint position={{ x: 1000, y: 240 }} isStop={stop1} />
      <StopPoint position={{ x: 880, y: 380 }} isStop={stop1} />
      <StopPoint position={{ x: 940, y: 380 }} isStop={stop1} />
      <StopPoint position={{ x: 1000, y: 380 }} isStop={stop1} />

      <SwitchPoint
        position={{ x: 160, y: 20 }}
        fromAngle={180}
        leftOutAngle={0}
        rightOutAngle={45}
        isLeft={switch1}
      />
      <SwitchPoint
        position={{ x: 320, y: 120 }}
        fromAngle={0}
        leftOutAngle={180}
        rightOutAngle={225}
        isLeft={switch1}
      />
      <SwitchPoint
        position={{ x: 380, y: 20 }}
        fromAngle={180}
        leftOutAngle={0}
        rightOutAngle={45}
        isLeft={switch1}
      />
      <SwitchPoint
        position={{ x: 440, y: 60 }}
        fromAngle={180}
        leftOutAngle={0}
        rightOutAngle={315}
        isLeft={switch1}
      />
      <SwitchPoint
        position={{ x: 420, y: 160 }}
        fromAngle={0}
        leftOutAngle={180}
        rightOutAngle={225}
        isLeft={switch1}
      />
      <SwitchPoint
        position={{ x: 480, y: 120 }}
        fromAngle={0}
        leftOutAngle={180}
        rightOutAngle={135}
        isLeft={switch1}
      />
      <SwitchPoint
        position={{ x: 600, y: 120 }}
        fromAngle={0}
        leftOutAngle={180}
        rightOutAngle={135}
        isLeft={switch1}
      />
      <SwitchPoint
        position={{ x: 800, y: 20 }}
        fromAngle={180}
        leftOutAngle={0}
        rightOutAngle={45}
        isLeft={switch1}
      />
      <SwitchPoint
        position={{ x: 840, y: 80 }}
        fromAngle={0}
        leftOutAngle={180}
        rightOutAngle={135}
        isLeft={switch1}
      />

      <SwitchPoint
        position={{ x: 1000, y: 280 }}
        fromAngle={270}
        leftOutAngle={90}
        rightOutAngle={135}
        isLeft={switch1}
      />
      <SwitchPoint
        position={{ x: 940, y: 480 }}
        fromAngle={90}
        leftOutAngle={270}
        rightOutAngle={225}
        isLeft={switch1}
      />
    </svg>
  )
}

export default RailroadMap
