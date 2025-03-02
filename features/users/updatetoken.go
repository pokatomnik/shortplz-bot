package users

import (
	userpkg "github.com/pokatomnik/shortplz-bot/entities/user"
	"github.com/samber/mo"
)

func (users Users) UpdateToken(userId int64, token string) mo.Result[struct{}] {
	result := users.db.
		Where(&userpkg.User{TelegramUserID: userId}).
		Updates(&userpkg.User{APIToken: token})
	if result.Error != nil {
		return mo.Err[struct{}](result.Error)
	}
	return mo.Ok(struct{}{})
}
