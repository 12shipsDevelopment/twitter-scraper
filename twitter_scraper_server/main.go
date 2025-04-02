package main

import (
	"fmt"
	"twitter_scraper_server/config"
	"twitter_scraper_server/handlers"
	"twitter_scraper_server/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.LoadConfigs()

	router := gin.Default()

	as := services.NewAccountService(cfg)

	handlers.HandleAccountRoutes(router, cfg, as)
	handlers.HandleTweetsRoutes(router, cfg, as)

	logrus.Info("Server started on :", cfg.Port)

	router.RunTLS(fmt.Sprintf(":%d", cfg.Port), cfg.CrtFile, cfg.KeyFile)
}
