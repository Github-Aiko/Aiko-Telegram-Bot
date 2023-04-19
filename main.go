package main

import (
	"fmt"

	"github.com/Github-Aiko/Aiko-Telegram-Bot/bot"
	"github.com/Github-Aiko/Aiko-Telegram-Bot/config"
	"github.com/Github-Aiko/Aiko-Telegram-Bot/data"
)

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetConfig().GetString("app.database.user"),
		config.GetConfig().GetString("app.database.pass"),
		config.GetConfig().GetString("app.database.ip"),
		config.GetConfig().GetInt("app.database.port"),
		config.GetConfig().GetString("app.database.name"),
	)

	data.New(dsn)

	tgbot, err := bot.New(config.GetConfig().GetString("bot.token"))
	if err != nil {
		panic(err)
	}

	bot.StartBot(tgbot)

}
