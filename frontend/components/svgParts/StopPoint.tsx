import React, { FC } from "react"
import { Point } from "../../types/svgPartsTypes"

interface Props {
  position: Point
  isStop: boolean
}

const StopPoint: FC<Props> = ({ position, isStop }) => (
  <g>
    <line
      x1={position.x}
      y1={position.y - 5}
      x2={position.x}
      y2={position.y + 5}
      stroke="black"
      strokeWidth={20}
      strokeLinecap="round"
    />
    <circle
      cx={position.x}
      cy={position.y - 6}
      r={5}
      fill={isStop ? "grey" : "green"}
    />
    <circle
      cx={position.x}
      cy={position.y + 6}
      r={5}
      fill={isStop ? "red" : "grey"}
    />
  </g>
)

export default StopPoint
