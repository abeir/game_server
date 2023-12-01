package testcase

import (
	"fmt"
	"game_server/server/common/conf"
	"game_server/server/common/utils"
	"os"
	"testing"
)

func TestEnvConfigFinder(t *testing.T) {
	pwd, _ := os.Getwd()
	full := fmt.Sprintf("%s/conf/test_server.yaml", pwd)
	_ = os.Setenv(conf.ServerConfigKey, full)

	finder := conf.EnvConfigFinder{}
	f, ok := finder.Find()
	if !ok {
		t.Errorf("not found config by env: %s", full)
	}
	exists, _ := utils.PathExists(f)
	if !exists {
		t.Errorf("config not exists by path: %s", f)
	}
}
