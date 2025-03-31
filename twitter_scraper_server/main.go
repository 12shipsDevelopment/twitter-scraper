package main

import (
	"fmt"
	"twitter_scraper_server/config"
	"twitter_scraper_server/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfigs()

	router := gin.Default()

	handlers.HandleAccountRoutes(router, cfg)
	handlers.HandleTweetsRoutes(router, cfg)

	fmt.Println("Server started on :", cfg.Port)
	router.Run(fmt.Sprintf(":%d", cfg.Port))
}
