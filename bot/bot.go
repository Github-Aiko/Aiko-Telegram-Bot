package bot

import (
	"time"

	"github.com/Github-Aiko/Aiko-Telegram-Bot/data"
	tele "gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

type Bot struct {
	*tele.Bot
}

func New(token string) (*Bot, error) {
	b, err := tele.NewBot(tele.Settings{
		URL:    "https://proxy-telegram-api.bps.im",
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return nil, err
	}

	return &Bot{
		Bot: b,
	}, nil
}

func StartBot(b *Bot, db *gorm.DB) {

	ur := data.NewUserRepo(db)

	service := NewBotService(ur)

	handle := NewBotHandle(service)

	handle.SetHandle(b)

	b.Start()
}
