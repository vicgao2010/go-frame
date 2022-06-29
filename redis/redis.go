package redis

import (
	client "github.com/go-redis/redis/v8"
	"github.com/vicgao-hub/go-frame/config"
	"github.com/vicgao-hub/go-frame/helper"
	"go.uber.org/zap"
)

func New(cfg *config.Config, logger *zap.Logger) (*client.Client, func(), error) {
	server := helper.SetDefaultString(cfg.Redis.Server, "127.0.0.1")
	db := helper.SetDefaultInt(cfg.Redis.DataBase, 0)
	password := helper.SetDefaultString(cfg.Redis.PassWord, "")
	pool := helper.SetDefaultInt(cfg.Redis.PoolSize, 10)
	redis := client.NewClient(&client.Options{
		Addr:     server,
		Password: password,
		DB:       db,
		PoolSize: pool,
	})
	cleanup := func() {
		if err := redis.Close(); err != nil {
			logger.Sugar().Errorf("redis close error %s", err)
		}
	}
	return redis, cleanup, nil
}
