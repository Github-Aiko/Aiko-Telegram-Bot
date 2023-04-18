package bot

import "net/http"

type Bot struct {
	Token  string
	Client *http.Client
}
