package database

import (
	"github.com/vicgao-hub/go-frame/config"
	"github.com/vicgao-hub/go-frame/helper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func New(cfg *config.Config, logger *zap.Logger) (*gorm.DB, func(), error) {
	prefix := helper.SetDefaultString(cfg.Database.Prefix, "")
	logger = logger.WithOptions(zap.AddCallerSkip(3))
	level := log.Silent
	if cfg.App.Debug {
		level = log.Info
	}
	db, err := gorm.Open(mysql.Open(cfg.Database.Dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: true,
		},
		Logger: NewLogger(logger).LogMode(level),
	})
	if err != nil {
		logger.Sugar().Errorf("db open error %s", err)
	}
	Db, err := db.DB()
	Db.SetMaxIdleConns(helper.SetDefaultInt(cfg.Database.MaxIdle, 5))
	Db.SetMaxOpenConns(helper.SetDefaultInt(cfg.Database.MaxOpen, 10))
	Db.SetConnMaxLifetime(helper.SetDefaultDuration(cfg.Database.LifeTime, "1m"))
	Db.SetConnMaxIdleTime(helper.SetDefaultDuration(cfg.Database.IdleTime, "2m"))
	cleanup := func() {
		if err := Db.Close(); err != nil {
			logger.Sugar().Errorf("db close error %s", err)
		}
	}
	return db, cleanup, err
}
