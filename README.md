# MiraiGo-module-random-mute

ID: `com.aimerneige.random.mute`

Module for [MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template)

## 功能

- 在群聊中接受到指定消息时对消息发送者进行 0 ~ 2 小时的禁言。

## 使用方法

在适当位置引用本包

```go
package example

imports (
    // ...

    _ "github.com/yukichan-bot-module/MiraiGo-module-random-mute"

    // ...
)

// ...
```

如果需要自定义关键词，在你的 `application.yaml` 里填入配置（支持多个关键词）：

```yaml
aimerneige:
  random_mute:
    keywords:
      - ".ban.me"
      - "随机禁言"
```

未填写配置文件时，关键词为默认关键词 `.ban.me` 和 `口球大礼包`，读取到配置后会覆盖默认关键词。
