import React, { FC } from "react"
import { Point } from "../../types/svgPartsTypes"
import { degToRad, radToDeg } from "../../utils/util"

interface Props {
  position: Point
  fromAngle: number
  leftOutAngle: number
  rightOutAngle: number
  isLeft: boolean
}

const SwitchPoint: FC<Props> = ({
  position,
  fromAngle,
  leftOutAngle,
  rightOutAngle,
  isLeft,
}) => {
  const fromAnglePointX = position.x + Math.cos(degToRad(fromAngle)) * 10
  const fromAnglePointY = position.y + Math.sin(degToRad(fromAngle)) * 10
  const leftOutAnglePointX = position.x + Math.cos(degToRad(leftOutAngle)) * 10
  const leftOutAnglePointY = position.y + Math.sin(degToRad(leftOutAngle)) * 10
  const rightOutAnglePointX =
    position.x + Math.cos(degToRad(rightOutAngle)) * 10
  const rightOutAnglePointY =
    position.y + Math.sin(degToRad(rightOutAngle)) * 10
  return (
    <g>
      <ellipse
        cx={position.x}
        cy={position.y}
        rx={10}
        fill="white"
        stroke="black"
      />
      <line
        x1={position.x}
        y1={position.y}
        x2={fromAnglePointX}
        y2={fromAnglePointY}
        stroke="black"
      />
      {isLeft ? (
        <line
          x1={position.x}
          y1={position.y}
          x2={leftOutAnglePointX}
          y2={leftOutAnglePointY}
          stroke="black"
        />
      ) : (
        <line
          x1={position.x}
          y1={position.y}
          x2={rightOutAnglePointX}
          y2={rightOutAnglePointY}
          stroke="black"
        />
      )}
    </g>
  )
}

export default SwitchPoint
