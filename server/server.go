package server

import (
	"github.com/vicgao-hub/go-frame/config"
	"github.com/vicgao-hub/go-frame/helper"
	"net/http"
)


func New(cfg *config.Config) *http.Server {
	addr := helper.SetDefaultString(cfg.Server.Addr, ":8080")
	timeout := helper.SetDefaultDuration(cfg.Server.Timeout, "1s")
	return &http.Server{Addr: addr, ReadTimeout: timeout, WriteTimeout: timeout}
}