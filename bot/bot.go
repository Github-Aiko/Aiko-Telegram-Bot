package bot

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func New(token string) *Bot {
	client := &http.Client{Timeout: 30 * time.Second}

	return &Bot{Token: token, Client: client}
}

func (b *Bot) SendMessage(chatID, text string) error {
	apiURL := fmt.Sprintf("https://proxy-telegram-api.bps.im/bot%s/sendMessage", b.Token)

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

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	return nil
}
