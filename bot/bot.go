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

func (b *Bot) request(url string, values url.Values) (data []byte, err error) {
	if b.Token == "" {
		return nil, errors.New("token cannot be empty")
	}

	apiURL := fmt.Sprintf("https://proxy-telegram-api.bps.im/bot%s%s?", b.Token, url)

	req, err := http.NewRequest("GET", apiURL+values.Encode(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(req)

	resp, err := b.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
