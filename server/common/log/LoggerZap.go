package log

import (
	"game_server/server/common/conf"
	"game_server/server/common/constant"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerZap struct {
	lv  LoggerLevel
	log *zap.SugaredLogger
}

func NewLoggerZap(config *conf.LoggerConfig) (ILogger, error) {
	level := GetLevel(config.Level)

	log := &LoggerZap{
		lv: level,
	}

	coreList := make([]zapcore.Core, 0, 2)
	for _, output := range config.Output {
		if strings.EqualFold(output, "console") {
			coreList = append(coreList, log.createConsoleCore(config, level))
		} else if strings.EqualFold(output, "file") {
			coreList = append(coreList, log.createFileCore(config, level))
		}
	}
	core := zapcore.NewTee(coreList...)
	logger := zap.New(core, zap.AddCaller())

	log.log = logger.Sugar()
	return log, nil
}

func (l *LoggerZap) SetLevel(lv LoggerLevel) {
	if lv >= LoggerDebug && lv <= LoggerError {
		l.lv = lv
	}
}

func (l *LoggerZap) GetLevel() LoggerLevel {
	return l.lv
}

func (l *LoggerZap) IsDebugEnabled() bool {
	return l.lv <= LoggerDebug
}

func (l *LoggerZap) IsInfoEnabled() bool {
	return l.lv <= LoggerInfo
}

func (l *LoggerZap) IsWarnEnabled() bool {
	return l.lv <= LoggerWarn
}

func (l *LoggerZap) IsErrorEnabled() bool {
	return l.lv <= LoggerError
}

func (l *LoggerZap) Debug(msg ...interface{}) {
	l.log.Debug(msg...)
}

func (l *LoggerZap) Info(msg ...interface{}) {
	l.log.Info(msg...)
}

func (l *LoggerZap) Warn(msg ...interface{}) {
	l.log.Warn(msg...)
}

func (l *LoggerZap) Error(msg ...interface{}) {
	l.log.Error(msg...)
}

func (l *LoggerZap) Debugf(fmt string, args ...interface{}) {
	l.log.Debugf(fmt, args...)
}

func (l *LoggerZap) Infof(fmt string, args ...interface{}) {
	l.log.Infof(fmt, args...)
}

func (l *LoggerZap) Warnf(fmt string, args ...interface{}) {
	l.log.Warnf(fmt, args...)
}

func (l *LoggerZap) Errorf(fmt string, args ...interface{}) {
	l.log.Errorf(fmt, args...)
}

func (l *LoggerZap) createConsoleCore(config *conf.LoggerConfig, lv LoggerLevel) zapcore.Core {
	fmt := getDateFormat(config.Console.DateFormat)
	encoderConf := zap.NewDevelopmentEncoderConfig()
	encoderConf.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(fmt))
	}
	encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConf.EncodeCaller = nil
	encoderConf.ConsoleSeparator = "  "

	encoder := zapcore.NewConsoleEncoder(encoderConf)
	writer := zapcore.Lock(os.Stdout)
	level := toZapLevel(lv)

	return zapcore.NewCore(encoder, writer, level)
}

func (l *LoggerZap) createFileCore(config *conf.LoggerConfig, lv LoggerLevel) zapcore.Core {
	fmt := getDateFormat(config.Console.DateFormat)
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(fmt))
		},
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		ConsoleSeparator: " - ",
	})

	writer := l.fileWriteSyncer(&config.File)
	level := toZapLevel(lv)

	return zapcore.NewCore(encoder, writer, level)
}

func (l *LoggerZap) fileWriteSyncer(config *conf.LoggerFile) zapcore.WriteSyncer {
	lumberjackLog := &lumberjack.Logger{
		Filename:   config.Filename,
		MaxSize:    config.MaxSize,
		MaxAge:     config.KeepDays,
		MaxBackups: config.MaxCount,
		Compress:   config.Compress,
	}
	return zapcore.AddSync(lumberjackLog)
}

func toZapLevel(lv LoggerLevel) zapcore.Level {
	switch lv {
	case LoggerDebug:
		return zapcore.DebugLevel
	case LoggerInfo:
		return zapcore.InfoLevel
	case LoggerWarn:
		return zapcore.WarnLevel
	case LoggerError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func getDateFormat(fmt string) string {
	if fmt == "" {
		return constant.TimestampFormat
	}
	return fmt
}

// encodeLevel 自定义日志级别显示
func encodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}
