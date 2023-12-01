package conf

import "os"

const ServerConfigKey = "GAME_SERVER_CONF"

type EnvConfigFinder struct {
}

func (e *EnvConfigFinder) Find() (string, bool) {
	return os.LookupEnv(ServerConfigKey)
}
