# ESP32に送る制御用JSONの仕様

## サーボモータ制御用

### 形式

以下のような JSON 形式で送信します。

```json
{
  "state": "<STATE>",
  "pin": "<PIN>",
  "angle": "<ANGLE>"
}
```

### `state`

送信後に期待する状態を送信します。以下の文字列を取ります。これ以外の文字列が送信されたときは失敗扱いとしてください。

| `state` | 説明 |
|---|---|
|`"ON"` |制御装置をON状態にする。|
|`"OFF"`|制御装置をOFF状態とする。|
|`"ANGLE"`|angleフィールドを見て角度を決めてください。|

制御が成功したならば HTTP Response Status Codeの200番台を送信してください。失敗したならばHTTP Response Status Codeの400番台を返してください。

### `angle`

stateと共に送信される場合があります。このフィールドは度数法による角度を指定します。有効なのは0度から359度です。数値が指定されます。

### `pin`

操作すべきピン番号情報です。情報はJSONの数値型で送信されます。もしそのピンを操作できないときはHTTP Response Status
Code [404 Not Found](https://developer.mozilla.org/ja/docs/Web/HTTP/Status/404) を返答してください。

# client <-> control external

クライアントからexternal、externalからクライアントへは次のような形式のjsonをweb socketでやりとりします。

クライアントは操作をしたときにexternalにデータを送信します。

externalは何らかの変化があったときや定期的に情報を提供するためにクライアントに向けてデータを送信します。

```json
{
  "station_name": "chofu_b1",
  "state": "ON"
}

```

station_nameは以下のgRPCのprotoファイルで定義されているものを使います。

https://github.com/ueckoken/plarail2021-soft/blob/main/backend/proto/statesync.proto

ONはストップレールの場合はレールを上げる(停止させる)状態です。
分岐レールの場合はOFFがまっすぐの方向、ONがまっすぐでない方向です。
