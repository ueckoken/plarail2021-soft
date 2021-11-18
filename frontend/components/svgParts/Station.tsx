import React, { FC } from "react"
import { Point } from "../../types/svgPartsTypes"

interface Props {
  name: string
  position: Point
}

const Station: FC<Props> = ({ name, position }) => {
  return (
    <g>
      <rect
        x={position.x - 50}
        y={position.y - 10}
        width={100}
        height={20}
        fill="white"
        stroke="black"
      />
      <text
        x={position.x}
        y={position.y}
        width={100}
        height={20}
        fontFamily="monospace"
        fontSize={20}
        textAnchor="middle"
        dominantBaseline="central"
      >
        {name}
      </text>
    </g>
  )
}

export default Station
