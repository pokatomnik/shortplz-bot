package env

import (
	"os"

	"github.com/samber/mo"
)

func (env *Env) GetBotToken() mo.Option[string] {
	if token, cached := env.data[botTokenEnvKey]; cached {
		return mo.Some(token)
	}
	if token, exists := os.LookupEnv(botTokenEnvKey); exists {
		env.data[botTokenEnvKey] = token
		return mo.Some(token)
	}
	return mo.None[string]()
}
