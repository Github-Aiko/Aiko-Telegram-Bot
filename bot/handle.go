package bot

type BotHandle struct {
	s *BotService
}

func NewBotHandle(s *BotService) *BotHandle {
	return &BotHandle{s: s}
}

func (s *BotService) SetHandle(b *Bot) {
	b.Handle("/start", s.startCmd)
	b.Handle("/ping", s.pingCmd)
	b.Handle("/login", s.loginCmd)
	b.Handle("/reg", s.registerCmd)
}
