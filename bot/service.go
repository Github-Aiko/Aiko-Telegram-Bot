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

// LoginCmd function
func (s *BotService) loginCmd(c tele.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return c.Reply("Usage: /login username password\n Usage: /login <username> <password>")
	}

	username := args[0]
	password := args[1]
	passwordMD5 := utlis.MD5(password)

	u, err := s.u.Login(username, passwordMD5)
	if err != nil {
		return c.Reply("Login failed: " + err.Error() + "\n Login failed")
	}

	return c.Reply("Login success\n Logged in successfully: " + u.Token)
}

// RegisterCmd function
func (s *BotService) registerCmd(c tele.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return c.Reply("Usage: /reg username password\n 使用方法: /reg 用户名 密码")
	}

	username := args[0]

	// validate that the username is at least 6 characters long
	if len(username) < 6 {
		return c.Reply("Username must be greater than 6 characters\n Username must be at least 6 characters long")
	}

	password := args[1]

	// validate that the password is at least 6 characters long
	if len(password) < 6 {
		return c.Reply("Password must be greater than 6 characters\n Password must be at least 6 characters long")
	}

	passwordMD5 := utlis.MD5(password)
	telegramID := c.Sender().ID
	token := utlis.UUID()

	u, err := s.u.CreateUser(username, passwordMD5, telegramID, token)
	if err != nil {
		return c.Reply("Register failed: " + err.Error() + "\n Registration failed")
	}

	return c.Reply("Register success\n Registered successfully: " + u.Token)
}
