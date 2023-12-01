package testcase

import (
	"game_server/server/common/conf"
	"game_server/test/common"
	"testing"
)

func TestYamlConfigLoader(t *testing.T) {
	finder := common.ConfigFinderImpl{}
	loader := conf.YamlConfigLoader{}
	loader.SetFinder(&finder)

	c, err := loader.Load()
	if err != nil {
		t.Error(err)
	}
	if c.Server.Listen != "0.0.0.0" {
		t.Errorf("Server.Listen expect: 0.0.0.0, but: %s", c.Server.Listen)
	}
	if c.Database.Uri == "" {
		t.Error("Database.Uri is empty")
	}
	if c.Logger.Level != "debug" {
		t.Errorf("Logger.Level expect: debug, but: %s", c.Logger.Level)
	}
}
