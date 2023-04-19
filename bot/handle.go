package bot

func setHandler(b *Bot) {
	b.Handle("/ping", ping)
}
