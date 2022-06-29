package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/vicgao-hub/go-frame/config"
	"github.com/vicgao-hub/go-frame/helper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	_ "path/filepath"
	_ "strconv"
)


func New(cfg *config.Config) (*zap.Logger, error) {
	level := filterZapAtomicLevel(cfg.Logger.Level)
	systemLevelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})
	return zap.New(zapcore.NewTee(
		zapcore.NewCore(getEncoder(cfg), getLogWriter(cfg), systemLevelEnabler)),
		zap.AddCaller(), zap.AddCallerSkip(0)), nil
}

func getEncoder(cfg *config.Config) zapcore.Encoder {
	var (
		durationEncoder = new(zapcore.DurationEncoder)
		callerEncoder   = new(zapcore.CallerEncoder)
		nameEncoder     = new(zapcore.NameEncoder)
		levelEncoder    = new(zapcore.LevelEncoder)
	)

	_ = durationEncoder.UnmarshalText([]byte(helper.SetDefaultString(cfg.Logger.Encoder.EncodeDuration, "ISO8601")))
	_ = callerEncoder.UnmarshalText([]byte(helper.SetDefaultString(cfg.Logger.Encoder.EncodeCaller, "full")))
	_ = nameEncoder.UnmarshalText([]byte(helper.SetDefaultString(cfg.Logger.Encoder.EncodeName, "full")))
	_ = levelEncoder.UnmarshalText([]byte(helper.SetDefaultString(cfg.Logger.Encoder.EncodeLevel, "capital")))

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        helper.SetDefaultString(cfg.Logger.Encoder.TimeKey, "ts"),
		LevelKey:       helper.SetDefaultString(cfg.Logger.Encoder.LevelKey, "level"),
		NameKey:        helper.SetDefaultString(cfg.Logger.Encoder.NameKey, "logger"),
		CallerKey:      helper.SetDefaultString(cfg.Logger.Encoder.CallerKey, "caller"),
		MessageKey:     helper.SetDefaultString(cfg.Logger.Encoder.MessageKey, "msg"),
		StacktraceKey:  helper.SetDefaultString(cfg.Logger.Encoder.StacktraceKey, "stacktrace"),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    *levelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: *durationEncoder,
		EncodeCaller:   *callerEncoder,
		EncodeName:     *nameEncoder,
	}
	var encoder zapcore.Encoder
	switch helper.SetDefaultString(cfg.Logger.Encoding, "json") {
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return encoder
}

func getLogWriter(cfg *config.Config) zapcore.WriteSyncer {
	if cfg.Logger.Path != "" {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   cfg.Logger.Path,
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   false,
		}
		return zapcore.AddSync(lumberJackLogger)
	}
	return zapcore.AddSync(os.Stdout)
}

func filterZapAtomicLevel(level string) zapcore.Level {
	var atomViper zapcore.Level
	switch level {
	default:
		atomViper = zap.InfoLevel
	case "debug":
		atomViper = zap.DebugLevel
	case "info":
		atomViper = zap.InfoLevel
	case "warn":
		atomViper = zap.WarnLevel
	case "error":
		atomViper = zap.ErrorLevel
	}
	return atomViper
}

