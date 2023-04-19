package bot

func setHandler(b *Bot) {
	b.Handle("/start", startCmd)
	b.Handle("/ping", pingCmd)
}
