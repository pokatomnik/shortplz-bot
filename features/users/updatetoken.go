package users

import (
	"strconv"

	"github.com/samber/mo"
	"go.mills.io/bitcask/v2"
)

func (users Users) UpdateToken(userId int64, token string) mo.Result[struct{}] {
	db, err := bitcask.Open(users.dbName)
	if err != nil {
		return mo.Err[struct{}](err)
	}
	defer db.Close()

	putErr := db.Put(bitcask.Key(strconv.Itoa(int(userId))), []byte(token))
	if putErr != nil {
		return mo.Err[struct{}](putErr)
	}

	return mo.Ok(struct{}{})
}
