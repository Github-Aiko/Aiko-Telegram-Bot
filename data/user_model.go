package data

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username   string `gorm:"type:varchar(50);not null"`
	Password   string `gorm:"type:varchar(50);not null"`
	TelegramID int64  `gorm:"type:bigint;not null"`
	Token      string `gorm:"type:varchar(50);not null"`
}

// GetTelegramID 通过 TelegramID 获取用户
// GetTelegramID get user by TelegramID
func GetTelegramID(TelegramID int64) (User, error) {
	var user User
	err := DB.Where("telegram_id = ?", TelegramID).First(&user).Error
	return user, err
}

// GetUsername 通过 Username 获取用户
// GetUsername get user by Username
func GetUsername(Username string) (User, error) {
	var user User
	err := DB.Where("username = ?", Username).First(&user).Error
	return user, err
}

// GetToken 通过 Token 获取用户
// GetToken get user by Token
func GetToken(Token string) (User, error) {
	var user User
	err := DB.Where("token = ?", Token).First(&user).Error
	return user, err
}

// CreateUser 创建用户
// CreateUser create user
func CreateUser(Username string, Password string, TelegramID int64, Token string) (User, error) {
	var user User
	user.Username = Username
	user.Password = Password
	user.TelegramID = TelegramID
	user.Token = Token
	err := DB.Create(&user).Error
	return user, err
}

// UpdateUser 更新用户
// UpdateUser update user
func UpdateUser(Username string, Password string, TelegramID int64, Token string) (User, error) {
	var user User
	user.Username = Username
	user.Password = Password
	user.TelegramID = TelegramID
	user.Token = Token
	err := DB.Save(&user).Error
	return user, err
}

// DeleteUser 删除用户
// DeleteUser delete user
func DeleteUser(Username string) error {
	var user User
	user.Username = Username
	err := DB.Delete(&user).Error
	return err
}

// Login 登录
// Login login
func Login(Username string, Password string) (User, error) {
	var user User
	err := DB.Where("username = ? AND password = ?", Username, Password).First(&user).Error
	return user, err
}
