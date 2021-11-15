# 2021 年調布祭プラレール企画配線図

## 配線図

[配線図 PDF](./chofufes2021-map.pdf)

![./chofufes2021-map.pdf](配線図)

駅付近にある`s`, `b` はそれぞれ制御するストップレール、分岐レールを表します。
設定ファイルなどで指定するときは以下のルールで指定します。

```text
<station_name>_<s or b><number>
```

例えば調布駅にあるストップレール `s1`を指定するときは `chofu_s1`と指定します。

サーボモータが ON, OFF になったときの挙動は [./api.md](https://github.com/ueckoken/plarail2021-soft/blob/main/docs/api.md#client---control-external)を参考にしてください。
