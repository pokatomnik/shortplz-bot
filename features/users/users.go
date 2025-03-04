package users

import (
	"github.com/samber/mo"
	"go.mills.io/bitcask/v2"
)

type Users struct {
	db *bitcask.Bitcask
}

func Open(dbName string) mo.Result[*Users] {
	db, err := bitcask.Open(dbName)
	if err != nil {
		return mo.Err[*Users](err)
	}

	users := &Users{db}

	return mo.Ok(users)
}

func (users *Users) Close() mo.Result[struct{}] {
	err := users.db.Close()
	if err != nil {
		return mo.Err[struct{}](err)
	}
	return mo.Ok(struct{}{})
}
