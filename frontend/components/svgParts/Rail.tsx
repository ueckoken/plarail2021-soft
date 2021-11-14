import React, { FC } from "react"
import { Point, TrainData } from "../../types/svgPartsTypes"

interface Props {
  startPosition: Point
  endPosition: Point
  trains: TrainData[]
}

const RailroadMap: FC<Props> = ({ startPosition, endPosition, trains }) => {
  return (
    <g>
      <line
        x1={startPosition.x}
        y1={startPosition.y}
        x2={endPosition.x}
        y2={endPosition.y}
        stroke="black"
      />
      {trains.map((train) => {
        const trainX =
          (endPosition.x - startPosition.x) * train.positionScale +
          startPosition.x
        const trainY =
          (endPosition.y - startPosition.y) * train.positionScale +
          startPosition.y
        return (
          <ellipse
            key={train.id}
            cx={trainX}
            cy={trainY}
            rx={7.5}
            fill="blue"
          />
        )
      })}
    </g>
  )
}

export default RailroadMap
