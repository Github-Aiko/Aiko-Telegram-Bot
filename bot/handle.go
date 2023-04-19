package bot

import "gopkg.in/telebot.v3/middleware"

type BotHandle struct {
	s *BotService
}

func NewBotHandle(s *BotService) *BotHandle {
	return &BotHandle{s: s}
}

func (s *BotService) SetHandle(b *Bot) {
	b.Use(middleware.Logger())
	b.Handle("/start", s.startCmd)
	b.Handle("/ping", s.pingCmd)
	b.Handle("/login", s.loginCmd)
	b.Handle("/reg", s.registerCmd)
}
