# 鯖エンドポイント

変更の可能性があります

## GKE

### 認証なし

- "control.chofufes2021.gotti.dev"
`./backend/external`が動いています。クライアントはここにむけてwebsocketを張ってください。
- "chofufes2021.gotti.dev"
`./frontend`が動いています。ここにメインのページがデプロイされます。
- "webrtc.chofufes2021.gotti.dev"
`./frontend/videoCast/one_to_multiple_cast_skyway.py`が動いています。webrtcのピアリングを行います。
- "auth.chofufes2021.gotti.dev"
認証画面です。認証が必要なページに入るには先にここを通ってください。
- "receiver-test.chofufes2021.gotti.dev"
`./frontend/videocast/skyway_receiver.html`が動いています。webrtcの受信側ページです。

### 認証あり

- "grafana.chofufes2021.gotti.dev"
grafanaというメトリクス可視化ツールが動いています。ID、パスワードはslackみてください。
- "prometheus.chofufes2021.gotti.dev"
prometheusというメトリクス収集ツールが動いています。基本的に見なくていいです。
- "alert.chofufes2021.gotti.dev"
使おうと思いましたがやめました。
- "webrtc-sender.chofufes2021.gotti.dev"
`./frontend/videocast/momo_sender.html`が動いています。webrtcの配信者側ページです。


## 学内

- "internal.chofufes2021.gotti.dev"
予定です。まだ取っていません。
- "speed.chofufes2021.gotti.dev"
予定です。まだ取っていません。
