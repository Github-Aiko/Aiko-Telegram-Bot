package bot

import (
	tele "gopkg.in/telebot.v3"
)

var (
	menu = &tele.ReplyMarkup{ResizeKeyboard: true}
)

func startCmd(c tele.Context) error {

	pingBtn := menu.Text("ping")

	menu.Reply(
		menu.Row(pingBtn),
	)

	c.Bot().Handle(&pingBtn, pingCmd)

	return c.Send("Hello!", menu)
}

func pingCmd(c tele.Context) error {
	return c.Send("pong!")
}
