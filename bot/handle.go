package bot

func setHandler(b *Bot) {
	b.Handle("/start", startCmd)
	b.Handle("/ping", pingCmd)
	b.Handle("/login", loginCmd)
	b.Handle("/reg", registerCmd)
}
