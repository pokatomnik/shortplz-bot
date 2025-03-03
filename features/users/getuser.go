package users

import (
	userpkg "github.com/pokatomnik/shortplz-bot/entities/user"
	"github.com/samber/mo"
)

func (users Users) GetUserOrCreate(userId int64) mo.Result[userpkg.User] {
	var user userpkg.User
	result := users.db.Where(&userpkg.User{TelegramUserID: userId}).FirstOrCreate(&user)

	if result.Error != nil {
		return mo.Err[userpkg.User](result.Error)
	}

	return mo.Ok(user)
}
