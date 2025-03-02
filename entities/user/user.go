package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TelegramUserID int64 `gorm:"uniqueIndex"`
	APIToken       string
}
