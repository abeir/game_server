package log

import "strings"

type LoggerLevel int

const (
	LoggerDebug LoggerLevel = iota
	LoggerInfo
	LoggerWarn
	LoggerError
)

type ILogger interface {
	SetLevel(lv LoggerLevel)
	GetLevel() LoggerLevel

	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool

	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})
	Debugf(fmt string, args ...interface{})
	Infof(fmt string, args ...interface{})
	Warnf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
}

func GetLevel(lv string) LoggerLevel {
	if lv == "" {
		return LoggerInfo
	}
	if strings.EqualFold(lv, "debug") {
		return LoggerDebug
	} else if strings.EqualFold(lv, "info") {
		return LoggerInfo
	} else if strings.EqualFold(lv, "warn") || strings.EqualFold(lv, "warning") {
		return LoggerWarn
	} else if strings.EqualFold(lv, "error") {
		return LoggerError
	}
	return LoggerInfo
}
