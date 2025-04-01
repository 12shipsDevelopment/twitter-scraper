package handlers

import (
	"net/http"
	"twitter_scraper_server/config"
	"twitter_scraper_server/services"

	"github.com/gin-gonic/gin"
)

func HandleAccountRoutes(router *gin.Engine, cfg *config.Config, as *services.AccountService) {
	// settings
	router.GET("1.1/account/settings.json", func(c *gin.Context) {
		settings, err := as.GetAccountSettings()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get 1.1/account/settings.json"})
			return
		}
		c.JSON(http.StatusOK, settings)
	})

	// list
	router.GET("1.1/account/multi/list.json", func(c *gin.Context) {
		list, err := as.GetAccountList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get 1.1/account/multi/list.json"})
			return
		}
		c.JSON(http.StatusOK, list)
	})

	// onboarding
	router.GET("1.1/onboarding/task.json", func(c *gin.Context) {
		var data map[string]interface{}
		err := c.BindJSON(&data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		flow, err := as.GetFlow(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get 1.1/onboarding/task.json"})
			return
		}
		c.JSON(http.StatusOK, flow)
	})

	// logout
	router.GET("1.1/account/logout.json", func(c *gin.Context) {
		as.Logout()
		c.JSON(http.StatusOK, "")
	})

	// GetAccessToken
	router.GET("oauth2/token", func(c *gin.Context) {
		token, err := as.GetAccessToken()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get 1.1/account/logout.json"})
			return
		}
		c.JSON(http.StatusOK, token)
	})
}
