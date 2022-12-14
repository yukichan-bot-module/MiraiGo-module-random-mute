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

## LICENSE

<a href="https://www.gnu.org/licenses/agpl-3.0.en.html">
<img src="https://www.gnu.org/graphics/agplv3-155x51.png">
</a>

本项目使用 `AGPLv3` 协议开源，您可以在 [GitHub](https://github.com/yukichan-bot-module/MiraiGo-module-random-mute) 获取本项目源代码。为了整个社区的良性发展，我们强烈建议您做到以下几点：

- **间接接触（包括但不限于使用 `Http API` 或 跨进程技术）到本项目的软件使用 `AGPLv3` 开源**
- **不鼓励，不支持一切商业使用**
