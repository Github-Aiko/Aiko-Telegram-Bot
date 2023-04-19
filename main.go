package main

import (
	"github.com/Github-Aiko/Aiko-Telegram-Bot/bot"
	"github.com/Github-Aiko/Aiko-Telegram-Bot/config"
)

func main() {
	// chatID := int64(123456789)
	// sendMessage(chatID, "Hello, world!")

	config := config.GetConfig()

	bot := bot.New(config.GetString("bot.token"))
	bot.SendMessage("-1001658662143", "Aiko 7")

}
