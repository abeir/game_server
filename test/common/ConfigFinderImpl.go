package common

type ConfigFinderImpl struct {
}

func (c *ConfigFinderImpl) Find() (string, bool) {
	return "D:/workspace/golang/game_server/test/conf/test_server.yaml", true
}
