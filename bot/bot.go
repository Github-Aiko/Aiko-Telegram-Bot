package bot

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func New(token string) (*Bot, error) {
	client := &http.Client{Timeout: 30 * time.Second}

	return &Bot{Token: token, Client: client}, nil
}

func (b *Bot) SendMessage(chatID, text string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Token)

	values := url.Values{}
	values.Set("chat_id", chatID)
	values.Set("text", text)

	resp, err := b.Client.Get(apiURL + "?" + values.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("failed to send message")
	}

	return nil
}
