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
func (s *BotService) startCmd(c tele.Context) error {

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	pingBtn := menu.Text("🏓 Ping")
	loginBtn := menu.Text("🔑 登录")
	registerBtn := menu.Text("📝 注册")

	menu.Reply(
		menu.Row(pingBtn),
		menu.Row(loginBtn),
		menu.Row(registerBtn),
	)

	c.Bot().Handle(&pingBtn, s.pingCmd)
	c.Bot().Handle(&loginBtn, s.loginCmd)
	c.Bot().Handle(&registerBtn, s.registerCmd)

	return c.Send("Hi! How can I help you today?", menu)
}

// PingCmd ping函数
func (s *BotService) pingCmd(c tele.Context) error {
	return c.Reply("🏓 pong!")
}

// LoginCmd 函数
func (s *BotService) loginCmd(c tele.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return c.Reply("用法: /login <用户名> <密码> 🔑\n示例: /login johnsmith password123 🔑")
	}

	username := args[0]
	password := args[1]
	passwordMD5 := utlis.MD5(password)

	u, err := s.u.Login(username, passwordMD5)
	if err != nil {
		return c.Reply("登录失败: " + err.Error() + " 🔒\n请使用正确的用户名和密码重试 🔑")
	}

	return c.Reply("登录成功 🎉\n您已成功登录，令牌为: " + u.Token + " ✅")
}

// RegisterCmd 函数
func (s *BotService) registerCmd(c tele.Context) error {
	args := c.Args()
	if len(args) != 2 {
		return c.Reply("用法: /reg <用户名> <密码> 📝\n示例: /reg johnsmith password123 📝")
	}

	username := args[0]

	// 验证用户名至少为6个字符长
	if len(username) < 6 {
		return c.Reply("用户名至少需要6个字符 🔒\n请使用至少6个字符的新用户名 📝")
	}

	password := args[1]

	// 验证密码至少为6个字符长
	if len(password) < 6 {
		return c.Reply("密码至少需要6个字符 🔒\n请使用至少6个字符的新密码 🔑")
	}

	passwordMD5 := utlis.MD5(password)
	telegramID := c.Sender().ID
	token := utlis.UUID()

	u, err := s.u.CreateUser(username, passwordMD5, telegramID, token)
	if err != nil {
		return c.Reply("注册失败: " + err.Error() + " 🔒\n请稍后重试 📝")
	}

	return c.Reply("注册成功 🎉\n您已成功注册，令牌为: " + u.Token + " ✅")
}
