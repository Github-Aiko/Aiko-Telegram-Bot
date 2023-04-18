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

func New(token, proxy string) (*Bot, error) {

	if proxy == "" {
		return &Bot{
			Token: token,
			Client: &http.Client{
				Timeout: 30 * time.Second,
			},
		}, nil
	}

	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	bot := &Bot{
		Token: token,
		Client: &http.Client{
			Transport: transport,
			Timeout:   30 * time.Second,
		},
	}

	return bot, nil
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
