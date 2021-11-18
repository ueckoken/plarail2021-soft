## Rail
車両存在区間は青で表示、後々表示用の車両コンポーネントを作るかも。
| プロパティ名 | 型 | 備考 |
| --- | --- | --- |
| positions | Point[] | 最低2つ |
| trains | number[] | 値は0.0-1.0の範囲 |

## Station
横向き白四角に駅名表示。
ストップレールが上がってる時はなんらかの表示。
駅をコンポーネント化するには複雑過ぎたので解体予定。
| プロパティ名 | 型 | 備考 |
| --- | --- | --- |
| name | String |  |
| position | Point |  |

## SwitchPoint
白円の中に入車角と出車角の線を表示。
| プロパティ名 | 型 | 備考 |
| --- | --- | --- |
| position | Point |  |
| fromAngle | number | 0.0-360.0 |
| leftOutAngle | number | 0.0-360.0 |
| rightOutAngle | number | 0.0-360.0 |
| isLeft | boolean | 切り替え |

## StopPoint
ストップレール用。
線路に重ねるのを想定。
| プロパティ名 | 型 | 備考 |
| --- | --- | --- |
| position | Point |  |
| isStop | boolean |  |