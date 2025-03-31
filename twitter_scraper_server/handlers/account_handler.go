package handlers

import (
	"net/http"
	"twitter_scraper_server/config"
	"twitter_scraper_server/services"

	"github.com/gin-gonic/gin"
)

func HandleAccountRoutes(router *gin.Engine, cfg *config.Config) {
	as := services.NewAccountService(cfg)

	// settings
	router.GET("1.1/account/settings.json", func(c *gin.Context) {
		settings, err := as.GetAccountSettings()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve address info"})
			return
		}
		c.JSON(http.StatusOK, settings)
	})

	// list
	router.GET("1.1/account/multi/list.json", func(c *gin.Context) {
		list, err := as.GetAccountList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve address info"})
			return
		}
		c.JSON(http.StatusOK, list)
	})
}
