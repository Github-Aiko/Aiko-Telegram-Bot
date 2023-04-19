package main

import (
	"fmt"

	"github.com/Github-Aiko/Aiko-Telegram-Bot/bot"
	"github.com/Github-Aiko/Aiko-Telegram-Bot/config"
	"github.com/Github-Aiko/Aiko-Telegram-Bot/data"
)

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetConfig().GetString("apps.database.user"),
		config.GetConfig().GetString("apps.database.pass"),
		config.GetConfig().GetString("apps.database.ip"),
		config.GetConfig().GetInt("apps.database.port"),
		config.GetConfig().GetString("apps.database.name"),
	)

	db, err := data.New(dsn)
	if err != nil {
		panic(err)
	}

	token := config.GetConfig().GetString("bot.token")
	fmt.Println(token)

	tgbot, err := bot.New(token)
	if err != nil {
		panic(err)
	}

	bot.StartBot(tgbot, db)

}
