## Rail
車両存在区間は青で表示、後々表示用の車両コンポーネントを作るかも。
| プロパティ名 | 型 | 備考 |
| --- | --- | --- |
| positions | Point[] | 最低2つ |
| trains | number[] | 値は0.0-1.0の範囲 |

## Platform
横向き白四角。
Rail、StopPointと組み合わせて駅を構成する。
| プロパティ名 | 型 | 備考 |
| --- | --- | --- |
| name | String | 駅名と番線 |
| position | Point |  |
| isHorizontal? | boolean | falseなら縦向き |
| size? | number | ホーム名がはみ出す場合に指定 |

## SwitchPoint
白円の中に入車角と出車角の線を表示。
アニメーションする。
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
見た目について再考が必要か。
| プロパティ名 | 型 | 備考 |
| --- | --- | --- |
| position | Point |  |
| isStop | boolean |  |