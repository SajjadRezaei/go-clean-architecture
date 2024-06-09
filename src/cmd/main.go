package main

import (
	"github.com/SajjadRezaei/go-clean-architecture/api"
	"github.com/SajjadRezaei/go-clean-architecture/config"
	"github.com/SajjadRezaei/go-clean-architecture/infra/cache"
	"github.com/SajjadRezaei/go-clean-architecture/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	api.InitServer(cfg)
}
