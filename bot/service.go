package bot

import (
	"github.com/Github-Aiko/Aiko-Telegram-Bot/data"
	"github.com/Github-Aiko/Aiko-Telegram-Bot/utlis"
	tele "gopkg.in/telebot.v3"
)

// StartCmd 启动函数
// StartCmd start function
func startCmd(c tele.Context) error {

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	pingBtn := menu.Text("ping")

	menu.Reply(
		menu.Row(pingBtn),
	)

	c.Bot().Handle(&pingBtn, pingCmd)

	return c.Send("Hello!", menu)
}

// PingCmd ping函数
// PingCmd ping function
func pingCmd(c tele.Context) error {
	return c.Send("pong!")
}

// LoginCmd 登录函数
// LoginCmd login function
func loginCmd(c tele.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return c.Reply("Usage: /login username password\n 使用方法: /login 用户名 密码")
	}

	username := args[0]
	password := args[1]
	passwordMD5 := utlis.MD5(password)

	u, err := data.Login(username, passwordMD5)
	if err != nil {
		return c.Reply("Login failed\n 登录失败")
	}

	return c.Reply("Login success\n 登入成功" + u.Token)
}

// RegisterCmd 注册函数
// RegisterCmd register function
func registerCmd(c tele.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return c.Reply("Usage: /login username password\n 使用方法: /login 用户名 密码")
	}

	username := args[0]

	//校验用户名是否不小于6位
	if len(username) < 6 {
		return c.Reply("Username must be greater than 6 characters\n 用户名必须大于6位")
	}

	password := args[1]

	//校验密码是否不小于6位
	if len(password) < 6 {
		return c.Reply("Password must be greater than 6 characters\n 密码必须大于6位")
	}

	passwordMD5 := utlis.MD5(password)
	telegramID := c.Sender().ID
	token := utlis.UUID()

	u, err := data.CreateUser(username, passwordMD5, telegramID, token)
	if err != nil {
		return c.Reply("Register failed\n 注册失败")
	}

	return c.Reply("Register success\n 注册成功\n" + u.Token)
}
