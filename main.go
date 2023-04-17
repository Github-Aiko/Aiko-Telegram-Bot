package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const botToken = "YOUR_BOT_TOKEN"

func sendMessage(chatID int64, text string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%d&text=%s", botToken, chatID, url.QueryEscape(text))
	http.Get(url)
}

func main() {
	chatID := int64(123456789)
	sendMessage(chatID, "Hello, world!")
}
