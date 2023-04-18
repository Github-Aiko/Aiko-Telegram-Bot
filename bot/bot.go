package bot

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Bot struct {
	Token  string
	Client *http.Client
}

func New(token, proxy string) *Bot {

	proxyUrl, _ := url.Parse(proxy)

	return &Bot{
		Token: token,
		Client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
			Timeout: 30 * time.Second,
		},
	}

}

func (b *Bot) SendMessage(chatID, text string) error {

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", b.Token, chatID, text)

	resp, err := b.Client.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("failed to send message")
	}

	return nil

}
