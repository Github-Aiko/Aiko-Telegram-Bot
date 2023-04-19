package bot

import (
	"github.com/Github-Aiko/Aiko-Telegram-Bot/data"
	"github.com/Github-Aiko/Aiko-Telegram-Bot/utlis"
	tele "gopkg.in/telebot.v3"
)

type BotService struct {
	u *data.UserRepo
}

func NewBotService(u *data.UserRepo) *BotService {
	return &BotService{u: u}
}

// StartCmd 启动函数
// StartCmd start function
func (s *BotService) startCmd(c tele.Context) error {

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	pingBtn := menu.Text("ping")
	loginBtn := menu.Text("login")
	registerBtn := menu.Text("register")

	menu.Reply(
		menu.Row(pingBtn),
		menu.Row(loginBtn),
		menu.Row(registerBtn),
	)

	c.Bot().Handle(&pingBtn, s.pingCmd)
	c.Bot().Handle(&loginBtn, s.loginCmd)
	c.Bot().Handle(&registerBtn, s.registerCmd)

	return c.Send("Hello!", menu)
}

// PingCmd ping函数
// PingCmd ping function
func (s *BotService) pingCmd(c tele.Context) error {
	return c.Send("pong!")
}

// LoginCmd 登录函数
// LoginCmd login function
func (s *BotService) loginCmd(c tele.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return c.Reply("Usage: /login username password\n 使用方法: /login 用户名 密码")
	}

	username := args[0]
	password := args[1]
	passwordMD5 := utlis.MD5(password)

	u, err := s.u.Login(username, passwordMD5)
	if err != nil {
		return c.Reply("Login failed\n 登录失败")
	}

	return c.Reply("Login success\n 登入成功" + u.Token)
}

// RegisterCmd 注册函数
// RegisterCmd register function
func (s *BotService) registerCmd(c tele.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return c.Reply("Usage: /reg username password\n 使用方法: /reg 用户名 密码")
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

	u, err := s.u.CreateUser(username, passwordMD5, telegramID, token)
	if err != nil {
		return c.Reply("Register failed\n 注册失败")
	}

	return c.Reply("Register success\n 注册成功\n" + u.Token)
}
