package bot

import (
	"time"

	tele "gopkg.in/telebot.v3"
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

func StartBot(b *Bot) {
	setHandler(b)
	b.Start()
}
