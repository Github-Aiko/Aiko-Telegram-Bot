package bot

import (
	"log"
	"net/url"
)

func (b *Bot) SendMessage(chatID, text string) error {

	values := url.Values{}
	values.Set("chat_id", chatID)
	values.Set("text", text)

	_, err := b.request("/sendMessage", values)
	if err != nil {
		log.Println(err)
	}

	return nil
}
