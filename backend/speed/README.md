# 仕様

これはスピードコントロールをするサーバです。ブラウザのクライアントからリクエストを受け取ってそれを車載のRaspberry Piに伝達します。

つまりブラウザに対してはサーバとして働き、Raspberry Piに対してはクライアントとして働きます。

例えば以下のような形式でデータを受けとれます。golangではhttpサーバを立ててQuery()で取得できます。(ぐぐって)

エンドポイントが決まってないので適当に決めてください。

https://speedcontrol.gotti.dev?trainId=takao&speed=8

このデータを事前定義の構造体にまとめてgRPCでRaspberry Piに送信します。

gRPCの定義は`../proto/speedControl.proto`にあります。
