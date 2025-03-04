package users

import (
	"strconv"

	"github.com/samber/mo"
	"go.mills.io/bitcask/v2"
)

func (users *Users) UpdateToken(userId int64, token string) mo.Result[struct{}] {
	putErr := users.db.Put(bitcask.Key(strconv.Itoa(int(userId))), []byte(token))
	if putErr != nil {
		return mo.Err[struct{}](putErr)
	}

	return mo.Ok(struct{}{})
}
