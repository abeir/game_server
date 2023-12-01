package testcase

import (
	"game_server/server/common/log"
	"game_server/server/common/utils"
	"game_server/test/common"
	"io"
	"os"
	"strings"
	"testing"
)

func TestLoggerZap(t *testing.T) {
	message := "test logger message"

	logger, logFile, err := common.NewLogger(true)
	if err != nil {
		t.Error(err)
		return
	}
	if logger.GetLevel() != log.LoggerDebug {
		t.Errorf("logger.GetLevel expect: %d, but: %d", log.LoggerDebug, logger.GetLevel())
	}
	logger.Info(message)

	file, err := os.OpenFile(logFile, os.O_RDONLY, 0644)
	if err != nil {
		t.Error(err)
		return
	}
	defer utils.CloseQuiet(file)

	data, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
		return
	}
	if !strings.Contains(string(data), message) {
		t.Error("log file has not contains test message")
	}
}
