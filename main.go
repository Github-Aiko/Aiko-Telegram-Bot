package main

import (
	"fmt"

	"github.com/Github-Aiko/Aiko-Telegram-Bot/config"
)

func main() {
	// chatID := int64(123456789)
	// sendMessage(chatID, "Hello, world!")
	config := config.GetConfig()
	fmt.Printf("Bot Telegram Token:%s", config.GetString("APPs.Telegram.Token"))
}
