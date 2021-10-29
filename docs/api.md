# ESP32に送る制御用JSONの仕様

## サーボモータ制御用

### 形式

以下のような JSON 形式で送信します。

```json
{
  "state": "<STATE>",
  "pin": "<PIN>"
}
```

### `state`

送信後に期待する状態を送信します。以下の文字列を取ります。

| `state` | 説明 |
|---|---|
|`"ON"` |制御装置をON状態にする。|
|`"OFF"`|制御装置をOFF状態とする。|

制御が成功したならば HTTP Response Status Codeの200番台を送信してください。失敗したならばHTTP Response Status Codeの400番台を返してください。

### `pin`

操作すべきピン番号情報です。情報はJSONの数値型で送信されます。もしそのピンを操作できないときはHTTP Response Status
Code [404 Not Found](https://developer.mozilla.org/ja/docs/Web/HTTP/Status/404) を返答してください。
