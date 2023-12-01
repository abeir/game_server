package conf

import (
	"time"
)

type ServerConfig struct {
	Server   Server       `yaml:"server"`
	Database Database     `yaml:"database"`
	Logger   LoggerConfig `yaml:"logger"`
}

type Server struct {
	Listen string `yaml:"listen"`
	Port   int    `yaml:"port"`
}

type Database struct {
	Uri         string        `yaml:"uri"`
	MaxLifetime time.Duration `yaml:"max_lifetime"`
	MaxIdleTime time.Duration `yaml:"max_idle_time"`
	MaxIdleConn int           `yaml:"max_idle_conn"`
	MaxOpenConn int           `yaml:"max_open_conn"`
}

type LoggerConfig struct {
	Level   string        `yaml:"level"`
	Output  []string      `yaml:"output"`
	Console LoggerConsole `yaml:"console"`
	File    LoggerFile    `yaml:"file"`
}

type LoggerConsole struct {
	DateFormat string `yaml:"date_format"`
}

type LoggerFile struct {
	DateFormat string `yaml:"date_format"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxCount   int    `yaml:"max_count"`
	KeepDays   int    `yaml:"keep_days"`
	Compress   bool   `yaml:"compress"`
}
