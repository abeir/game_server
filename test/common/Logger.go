package common

import (
	"fmt"
	"game_server/server/common/conf"
	"game_server/server/common/log"
	"os"
)

func NewLogger(remove bool) (logger log.ILogger, filePath string, err error) {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	filePath = fmt.Sprintf("%s/temp/game_server_test.log", pwd)
	if remove {
		_ = os.Remove(filePath)
	}

	config := conf.LoggerConfig{
		Level:  "debug",
		Output: []string{"console", "file"},
		Console: conf.LoggerConsole{
			DateFormat: "2006-01-02 15:04:05.000",
		},
		File: conf.LoggerFile{
			DateFormat: "2006-01-02 15:04:05.000",
			Filename:   filePath,
			MaxSize:    10,
			MaxCount:   2,
			KeepDays:   3,
			Compress:   true,
		},
	}
	logger, err = log.NewLoggerZap(&config)
	return
}
