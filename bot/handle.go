package bot

import (
	"gopkg.in/telebot.v3/middleware"

	tele "gopkg.in/telebot.v3"
)

type BotHandle struct {
	s *BotService
}

func NewBotHandle(s *BotService) *BotHandle {
	return &BotHandle{s: s}
}

func (h *BotHandle) SetHandle(b *Bot) {
	b.Use(middleware.Logger())
	b.Handle("/start", h.s.startCmd)
	b.Handle("/ping", h.s.pingCmd)
	b.Handle("/login", h.s.loginCmd)
	b.Handle("/reg", h.s.registerCmd)

	b.SetCommands([]tele.Command{
		{Text: "start", Description: "Start the bot"},

		{Text: "ping", Description: "Ping the bot"},

		{Text: "login", Description: "Login the bot"},
	})

}
