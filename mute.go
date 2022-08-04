package mute

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
)

var instance *mute
var logger = utils.GetModuleLogger("com.aimerneige.random.mute")
var keywordList []string

type mute struct {
}

func init() {
	instance = &mute{}
	bot.RegisterModule(instance)
}

func (m *mute) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "com.aimerneige.random.mute",
		Instance: instance,
	}
}

func (m *mute) Init() {
	configKeywordList := config.GlobalConfig.GetStringSlice("aimerneige.random_mute.keywords")
	if len(configKeywordList) == 0 {
		keywordList = []string{".ban.me", "口球大礼包"}
	} else {
		for _, k := range configKeywordList {
			keywordList = append(keywordList, k)
		}
	}
}

func (m *mute) PostInit() {
}

func (m *mute) Serve(b *bot.Bot) {
	b.OnGroupMessage(func(c *client.QQClient, msg *message.GroupMessage) {
		if isKeyword(msg.ToString()) {
			groupMemberInfo, err := c.GetMemberInfo(msg.GroupCode, msg.Sender.Uin)
			if err != nil {
				logger.WithError(err).Errorf("Fail to get group member info.")
				return
			}
			seed := rand.NewSource(time.Now().UnixNano())
			rand := rand.New(seed)
			muteDuration := rand.Intn(7200)
			groupMemberInfo.Mute(uint32(muteDuration))
			replyStr := fmt.Sprintf("恭喜%s获得了%d秒的禁言大礼包", msg.Sender.CardName, muteDuration)
			c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText(replyStr)))
		}
	})
}

func (m *mute) Start(b *bot.Bot) {
}

func (m *mute) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

func isKeyword(msg string) bool {
	for _, s := range keywordList {
		if s == msg {
			return true
		}
	}
	return false
}
