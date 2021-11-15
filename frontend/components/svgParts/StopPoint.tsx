import React, { FC } from "react"
import { Point } from "../../types/svgPartsTypes"

interface Props {
  position: Point
  isStop: boolean
}

const StopPoint: FC<Props> = ({ position, isStop }) => {
  if (!isStop) {
    return null
  }
  return (
    <g>
      <circle cx={position.x} cy={position.y} r={10} fill="red" />
      <rect
        x={position.x - 7}
        y={position.y - 2}
        width={14}
        height={4}
        fill="white"
      />
    </g>
  )
}

export default StopPoint
