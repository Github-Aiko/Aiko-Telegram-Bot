package bot

import (
	"fmt"
	"log"
	"net/url"
)

func (b *Bot) SendMessage(chatID, text string) error {

	values := url.Values{}
	values.Set("chat_id", chatID)
	values.Set("text", text)

	data, err := b.request("/sendMessage", values)
	if err != nil {
		log.Println(err)
	}

	fmt.Print(string(data))

	return nil
}
