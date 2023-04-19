package bot

import (
	tele "gopkg.in/telebot.v3"
)

func ping(c tele.Context) error {
	return c.Send("pong!")
}
