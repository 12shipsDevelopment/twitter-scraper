package handlers

import (
	"net/http"
	"twitter_scraper_server/config"
	"twitter_scraper_server/services"

	"github.com/gin-gonic/gin"
)

func HandleTweetsRoutes(router *gin.Engine, cfg *config.Config, as *services.AccountService) {
	ts := services.NewTweetsService(cfg, as)

	// UserTweetsAndReplies
	router.GET("/i/api/graphql/bt4TKuFz4T7Ckk-VvQVSow/UserTweetsAndReplies", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetUserTweetsAndReplies(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user tweets and replies"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// UserTweets
	router.GET("/i/api/graphql/UGi7tjRPr-d_U3bCPIko5Q/UserTweets", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetUserTweets(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user tweets"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// profile
	router.GET("/2/timeline/profile/{:userID}.json", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.FetchTweetsByUserIDLegacy(encodeQuery, c.Param("userID"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user tweets"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// conversation
	router.GET("/2/timeline/conversation/{:id}.json", func(c *gin.Context) {
		timeline, err := ts.GetTweet1(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch conversation"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// GetTweet2 - TweetDetail
	router.GET("/i/api/graphql/VWFGPVAGkZMGRKGe3GFFnA/TweetDetail", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetTweet2(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweet detail"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// TweetResultByRestId
	router.GET("/i/api/graphql/xBtHv5-Xsk268T5ng_OGNg/TweetResultByRestId", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetTweetResultByRestId(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweet result by restId"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// HomeLatestTimeline
	router.GET("/i/api/graphql/9EwYy8pLBOSFlEoSP2STiQ/HomeLatestTimeline", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetHomeLatestTimeline(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweet result by restId"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// HomeTimeline
	router.GET("/i/api/graphql/1u0Wlkw6Ru1NwBUD-pDiww/HomeTimeline", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetHomeTimeline(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweet result by restId"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// Bookmarks
	router.GET("/i/api/graphql/-IyJFt9_jS_9d_vS3NN-fA/Bookmarks", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetBookmarks(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookmarks"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// Following
	router.GET("/i/api/graphql/g5P4cbXR4ta4oCeE7y2vLQ/Following", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetFollowing(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch following"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// Followers
	router.GET("/i/api/graphql/jwbfbSzn0FRL_AMZGsYDag/Followers", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetFollowers(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch followers"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// UserMedia
	router.GET("/i/api/graphql/2tLOJWwGuCTytDrGBg8VwQ/UserMedia", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetUserMedia(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user media"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// UserByScreenName
	router.GET("/graphql/Yka-W8dz7RaEuQNkroPkYw/UserByScreenName", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetUserByScreenName(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user by username"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// UserByRestId
	router.GET("/i/api/graphql/Qw77dDjp9xCpUY-AXwt-yQ/UserByRestId", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		user, err := ts.GetUserByRestId(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user by restId"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	// TweetDetail
	router.GET("/i/api/graphql/ldqoq5MmFHN1FhMGvzC9Jg/TweetDetail", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		user, err := ts.GetTweetDetail(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweet detail"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	// FetchScheduledTweets
	router.GET("/i/api/graphql/ITtjAzvlZni2wWXwf295Qg/FetchScheduledTweets", func(c *gin.Context) {
		encodeQuery := c.Request.URL.Query().Encode()

		user, err := ts.GetFetchScheduledTweets(encodeQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch scheduled tweets"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	// DeleteScheduledTweet
	router.POST("/i/api/graphql/CTOVqej0JBXAZSwkp1US0g/DeleteScheduledTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response, err := ts.DeleteScheduledTweet(payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete scheduled tweet"})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	// CreateScheduledTweet
	router.POST("/i/api/graphql/LCVzRQGxOaGnOnYH01NQXg/CreateScheduledTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response, err := ts.CreateScheduledTweet(payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create scheduled tweet"})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	// SearchTimeline
	router.GET("/i/api/graphql/nK1dw4oV3k4w5TdtcAdSww/SearchTimeline", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		encodedQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.SearchTimeline(encodedQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search timeline"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// AudioSpaceById
	router.GET("/i/api/graphql/d03OdorPdZ_sH9V3D1_yWQ/AudioSpaceById", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		encodedQuery := c.Request.URL.Query().Encode()

		spaceData, err := ts.GetAudioSpaceById(encodedQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch audio space by id"})
			return
		}
		c.JSON(http.StatusOK, spaceData)
	})

	// guide
	router.GET("/2/guide.json", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		encodedQuery := c.Request.URL.Query().Encode()

		timeline, err := ts.GetGuide(encodedQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch guide"})
			return
		}
		c.JSON(http.StatusOK, timeline)
	})

	// CreateTweet
	router.POST("/i/api/graphql/oB-5XsHNAbjvARJEc8CZFw/CreateTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response, err := ts.CreateTweet(payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tweet"})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	// DeleteTweet
	router.POST("/i/api/graphql/VaenaVgh5q5ih7kvyVjgtg/DeleteTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response, err := ts.DeleteTweet(payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tweet"})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	// CreateRetweet
	router.POST("/i/api/graphql/ojPdsZsimiJrUGLR1sjUtA/CreateRetweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response, err := ts.CreateRetweet(payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create retweet"})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	// DeleteRetweet
	router.POST("/i/api/graphql/iQtK4dl5hBmXewYZuEOKVw/DeleteRetweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response, err := ts.DeleteRetweet(payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete retweet"})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	// FavoriteTweet
	router.POST("/i/api/graphql/lI07N6Otwv1PhnEgXILM7A/FavoriteTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response, err := ts.FavoriteTweet(payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to favorite tweet"})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	// UnfavoriteTweet
	router.POST("/i/api/graphql/ZYKSe-w7KEslx3JhSIk5LA/UnfavoriteTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response, err := ts.UnfavoriteTweet(payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfavorite tweet"})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	// GetRetweeters
	router.GET("/i/api/graphql/8019obfgnveiPiJuS2Rtow/Retweeters", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		response, err := ts.GetRetweeters(encodedQuery)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch retweeters"})
			return
		}
		c.JSON(http.StatusOK, response)
	})
}
