package data

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username   string `gorm:"type:varchar(50);not null"`
	Password   string `gorm:"type:varchar(50);not null"`
	TelegramID int64  `gorm:"type:bigint;not null"`
	Token      string `gorm:"type:varchar(50);not null"`
	IsAdmin    bool   `gorm:"type:tinyint(1);not null"`
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

// GetTelegramID 通过 TelegramID 获取用户
// GetTelegramID get user by TelegramID
func (u *UserRepo) GetTelegramID(TelegramID int64) (User, error) {
	var user User
	err := u.db.Where("telegram_id = ?", TelegramID).First(&user).Error
	return user, err
}

// GetUsername 通过 Username 获取用户
// GetUsername get user by Username
func (u *UserRepo) GetUsername(Username string) (User, error) {
	var user User
	err := u.db.Where("username = ?", Username).First(&user).Error
	return user, err
}

// GetToken 通过 Token 获取用户
// GetToken get user by Token
func (u *UserRepo) GetToken(Token string) (User, error) {
	var user User
	err := u.db.Where("token = ?", Token).First(&user).Error
	return user, err
}

// CreateUser 创建用户
// CreateUser create user
func (u *UserRepo) CreateUser(Username string, Password string, TelegramID int64, Token string) (User, error) {
	var user User
	user.Username = Username
	user.Password = Password
	user.TelegramID = TelegramID
	user.Token = Token
	err := u.db.Create(&user).Error
	return user, err
}

// UpdateUser 更新用户
// UpdateUser update user
func (u *UserRepo) UpdateUser(Username string, Password string, TelegramID int64, Token string) (User, error) {
	var user User
	user.Username = Username
	user.Password = Password
	user.TelegramID = TelegramID
	user.Token = Token
	err := u.db.Save(&user).Error
	return user, err
}

// DeleteUser 删除用户
// DeleteUser delete user
func (u *UserRepo) DeleteUser(Username string) error {
	var user User
	user.Username = Username
	err := u.db.Delete(&user).Error
	return err
}

// Login 登录
// Login login
func (u *UserRepo) Login(Username string, Password string) (User, error) {
	var user User
	err := u.db.Where("username = ? AND password = ?", Username, Password).First(&user).Error
	return user, err
}
