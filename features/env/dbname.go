package env

import (
	"os"

	"github.com/samber/mo"
)

func (env *Env) GetDBFileName() mo.Option[string] {
	if fname, cached := env.data[dbFileNameEnvKey]; cached {
		return mo.Some(fname)
	}
	if fname, exists := os.LookupEnv(dbFileNameEnvKey); exists {
		env.data[botTokenEnvKey] = fname
		return mo.Some(fname)
	}
	return mo.None[string]()
}
