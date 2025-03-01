package env

import (
	"os"

	"github.com/samber/mo"
)

func (env *Env) GetAPIToken() mo.Option[string] {
	if token, cached := env.data[apiTokenEnvKey]; cached {
		return mo.Some(token)
	}
	if token, exists := os.LookupEnv(apiTokenEnvKey); exists {
		env.data[apiTokenEnvKey] = token
		return mo.Some(token)
	}
	return mo.None[string]()
}
