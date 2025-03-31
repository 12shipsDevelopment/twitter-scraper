package handlers

import (
	"encoding/json"
	"net/http"
	"twitter_scraper_server/config"
	"twitter_scraper_server/services"

	"github.com/gin-gonic/gin"
)

func HandleTweetsRoutes(router *gin.Engine, cfg *config.Config) {
	ts := services.NewTweetsService(cfg)

	// UserTweetsAndReplies
	router.GET("i/api/graphql/bt4TKuFz4T7Ckk-VvQVSow/UserTweetsAndReplies", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	// UserTweets
	router.GET("/i/api/graphql/UGi7tjRPr-d_U3bCPIko5Q/UserTweets", func(c *gin.Context) {
		var variables map[string]interface{}
		err := json.Unmarshal([]byte(c.Query("variables")), &variables)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID := variables["userID"].(string)
		maxTweetsNbr := variables["maxTweetsNbr"].(int)
		cursor := variables["cursor"].(string)

		timeline, err := ts.GetUserTweets(userID, maxTweetsNbr, cursor)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve address info"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// profile
	router.GET("2/timeline/profile/{:userID}.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	// conversation
	router.GET("2/timeline/conversation/{:id}.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	// TweetDetail
	router.GET("i/api/graphql/VWFGPVAGkZMGRKGe3GFFnA/TweetDetail", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	// TweetResultByRestId
	router.GET("i/api/graphql/xBtHv5-Xsk268T5ng_OGNg/TweetResultByRestId", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	// HomeLatestTimeline
	router.GET("i/api/graphql/9EwYy8pLBOSFlEoSP2STiQ/HomeLatestTimeline", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})

	// HomeTimeline
	router.GET("i/api/graphql/1u0Wlkw6Ru1NwBUD-pDiww/HomeTimeline", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})
}
