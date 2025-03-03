package users

import (
	"errors"
	"strconv"

	"github.com/pokatomnik/shortplz-bot/entities/user"
	"github.com/samber/mo"
	"go.mills.io/bitcask/v2"
)

func (users Users) GetUser(userId int64) mo.Result[user.User] {
	db, err := bitcask.Open(users.dbName)
	if err != nil {
		return mo.Err[user.User](errors.New(errorOpenDatabase))
	}
	defer db.Close()

	v, err := db.Get(bitcask.Key(strconv.Itoa(int(userId))))
	if err != nil {
		return mo.Err[user.User](err)
	}

	token := string(v)

	return mo.Ok(user.User{TelegramUserID: userId, APIToken: token})
}
