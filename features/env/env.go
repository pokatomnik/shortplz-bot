package env

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Env struct {
	data map[string]string
}

func New() *Env {
	err := godotenv.Load()
	if err != nil {
		logrus.Warn("Missing .env file, proceed with OS env only")
	}
	return &Env{
		data: make(map[string]string),
	}
}
