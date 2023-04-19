package main

import (
	"github.com/Github-Aiko/Aiko-Telegram-Bot/bot"
	"github.com/Github-Aiko/Aiko-Telegram-Bot/config"
)

func main() {

	tgbot, err := bot.New(config.GetConfig().GetString("bot.token"))
	if err != nil {
		panic(err)
	}

	bot.Start(tgbot)

}
