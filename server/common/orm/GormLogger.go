package orm

import (
	"context"
	"fmt"
	"game_server/server/common/log"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type GormLogger struct {
	level         logger.LogLevel
	slowThreshold time.Duration

	log log.ILogger
}

func newGormLogger(l log.ILogger, slowThreshold time.Duration) *GormLogger {
	lv := logger.Warn
	switch l.GetLevel() {
	case log.LoggerError:
		lv = logger.Error
	case log.LoggerWarn:
		lv = logger.Warn
	case log.LoggerInfo:
		lv = logger.Warn
	case log.LoggerDebug:
		lv = logger.Info
	}
	return &GormLogger{
		level:         lv,
		slowThreshold: slowThreshold,
		log:           l,
	}
}

func (g *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	switch level {
	case logger.Silent:
		g.log.SetLevel(log.LoggerError)
	case logger.Error:
		g.log.SetLevel(log.LoggerError)
	case logger.Warn:
		g.log.SetLevel(log.LoggerWarn)
	case logger.Info:
		g.log.SetLevel(log.LoggerDebug)
	}
	g.level = level
	return g
}

func (g *GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	g.log.Debugf(s, i...)
}

func (g *GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	g.log.Warnf(s, i...)
}

func (g *GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	g.log.Errorf(s, i...)
}

var (
	traceStr     = "%s - [%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s - [%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s - [%.3fms] [rows:%v] %s"
)

func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && g.level >= logger.Error:
		sql, rows := fc()
		if rows == -1 {
			g.log.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			g.log.Errorf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > g.slowThreshold && g.slowThreshold != 0 && g.level >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", g.slowThreshold)
		if rows == -1 {
			g.log.Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			g.log.Warnf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case g.level == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			g.log.Debugf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			g.log.Debugf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
