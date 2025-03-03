package user

type User struct {
	TelegramUserID int64 `gorm:"uniqueIndex"`
	APIToken       string
}
