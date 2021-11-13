## 仕様

これはスピードコントロールをするサーバです。ブラウザのクライアントからリクエストを受け取ってそれを車載のRaspberry Piに伝達します。

つまりブラウザに対してはサーバとして働き、Raspberry Piに対してはクライアントとして働きます。

例えば以下のような形式でデータを受けとれます。golangではhttpサーバを立ててQuery()で取得できます。(ぐぐって)

エンドポイントは決まってないので適当に決めてください。たとえば

https://speedcontrol.gotti.dev?trainId=takao&speed=8

このデータを事前定義の構造体にまとめてgRPCでRaspberry Piに送信します。

gRPCの定義は`../proto/speedControl.proto`にあります。gRPCはこのへんでも使われてるので参考にしてください。

https://github.com/ueckoken/plarail2021-soft/blob/main/backend/external/pkg/servo/send.go
