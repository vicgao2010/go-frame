package database

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	log "gorm.io/gorm/logger"
	_ "path/filepath"
	_ "strconv"
	"time"
)

type Config struct {
	SlowThreshold             time.Duration
	Colorful                  bool
	IgnoreRecordNotFoundError bool
	Level                  log.LogLevel
}

type Logger struct {
	Config
	logger        *zap.Logger
}

func NewLogger(log *zap.Logger) log.Interface {
	return &Logger{logger:log}
}

func (l *Logger) LogMode(level log.LogLevel) log.Interface {
	clone := *l
	clone.Level = level
	return &clone
}

func (l Logger) Info(_ context.Context, msg string, args ...interface{}) {
	l.logger.Sugar().Debugf(msg, args...)
}

func (l Logger) Warn(_ context.Context, msg string, args ...interface{}) {
	l.logger.Sugar().Warnf(msg, args...)
}

func (l Logger) Error(_ context.Context, msg string, args ...interface{}) {
	l.logger.Sugar().Errorf(msg, args...)
}

func (l Logger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.Level <= log.Silent {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.Level >= log.Error && (!errors.Is(err, log.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.logger.Sugar().Error(err, float64(elapsed.Nanoseconds())/1e6, "-", " ", sql)
		} else {
			l.logger.Sugar().Error(err, float64(elapsed.Nanoseconds())/1e6, rows, " ", sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.Level >= log.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.logger.Sugar().Warn(slowLog, float64(elapsed.Nanoseconds())/1e6, "-", " ", sql)
		} else {
			l.logger.Sugar().Warn(slowLog, float64(elapsed.Nanoseconds())/1e6, rows, " ", sql)
		}
	case l.Level == log.Info:
		sql, rows := fc()
		if rows == -1 {
			l.logger.Sugar().Debug(float64(elapsed.Nanoseconds())/1e6, "-", " ", sql)
		} else {
			l.logger.Sugar().Debug(float64(elapsed.Nanoseconds())/1e6, rows, " ", sql)
		}
	}
}
