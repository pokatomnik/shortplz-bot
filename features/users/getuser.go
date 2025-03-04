package users

import (
	"errors"
	"strconv"

	"github.com/pokatomnik/shortplz-bot/entities/user"
	"github.com/samber/mo"
	"go.mills.io/bitcask/v2"
)

func (users *Users) GetUser(userId int64) mo.Result[user.User] {
	v, err := users.db.Get(bitcask.Key(strconv.Itoa(int(userId))))
	if err != nil {
		if errors.Is(err, bitcask.ErrKeyNotFound) {
			return mo.Ok(user.User{TelegramUserID: userId})
		} else {
			return mo.Err[user.User](err)
		}
	}

	token := string(v)

	return mo.Ok(user.User{TelegramUserID: userId, APIToken: token})
}
