# plarail2021-soft

## これはなに？

このリポジトリでは電気通信大学工学研究部 調布祭展示企画のためのソースコードを管理します。

## コントリビュータのみなさんへ

GitHub Actions による CI/CD を走らせるために main ブランチへの直接コミットを禁止しています。 プルリクエストを発行してください。

## 動かす方法

- [環境変数ドキュメント](./docs/environmentValList.md)に従って環境変数を設定する。

## CI/CDパイプラインについて

このリポジトリではCI/CDによるビルド、デプロイの自動化を行なっています。
github actionsによりイメージのビルドを行い、fluxcdにより新しいイメージを自動でサーバに適用しています。

### 使い方

どんどんmainにマージしていってください。./manifests以下への変更は即時でfluxcdにより反映されます。

dockerイメージの更新やるぞ!と思ったらmainからdeploymentにマージしてください。github actionsによるdocker imageのビルドが走ります。

新しいイメージができたらfluxcdが自動で検出してmainに更新commitを加えるところまでやります。鯖への反映は2分ぐらいかかります。

鯖エンドポイント(変更の可能性があります)

- backend/external (ポイント制御) https://plarail2021-backend.gotti.dev
- frontend (フロントエンド全般) https://plarail2021.gotti.dev
- frontend/other/one_to_multiple_cast_skyway.py (STUN鯖) https://plarail2021-py.gotti.dev

## 各種ドキュメント

- [環境変数に関するドキュメント](./docs/environmentValList.md)
- [ESP32に送信するJSONに関するドキュメント](./docs/api.md)
- [External<->Internalで用いるプロトコルに関するドキュメント(未整備)](./docs/protocolBuf.md)
