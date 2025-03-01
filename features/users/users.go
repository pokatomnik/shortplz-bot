package users

import (
	"errors"

	"github.com/pokatomnik/shortplz-bot/entities/user"
	"github.com/samber/mo"
	"gorm.io/gorm"
)

type Users struct {
	db *gorm.DB
}

func Open(dialector gorm.Dialector) mo.Result[Users] {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return mo.Err[Users](errors.New(errorOpenDatabase))
	}

	mErr := db.AutoMigrate(&user.User{})
	if mErr != nil {
		return mo.Err[Users](errors.New(errorMigrate))
	}

	return mo.Ok(Users{db})
}
