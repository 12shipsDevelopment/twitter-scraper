package handlers

import (
	"net/http"
	"twitter_scraper_server/config"

	"github.com/gin-gonic/gin"
)

func HandleAuthRoutes(router *gin.Engine, cfg *config.Config) {
	// as := services.NewAuthService(cfg)

	// onboarding
	router.GET("1.1/onboarding/task.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	// logout
	router.GET("1.1/account/logout.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	// oauth
	router.GET("oauth2/token", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})
}
