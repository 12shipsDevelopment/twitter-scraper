package handlers

import (
	"net/http"
	"twitter_scraper_server/config"
	"twitter_scraper_server/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleTweetsRoutes(router *gin.Engine, cfg *config.Config, as *services.AccountService) {
	ts := services.NewTweetsService(cfg)

	// UserTweetsAndReplies
	router.GET("/i/api/graphql/bt4TKuFz4T7Ckk-VvQVSow/UserTweetsAndReplies", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetUserTweetsAndReplies(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetUserTweetsAndReplies via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user tweets and replies"})
	})

	// UserTweets
	router.GET("/i/api/graphql/UGi7tjRPr-d_U3bCPIko5Q/UserTweets", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetUserTweets(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetUserTweets via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user tweets"})
	})

	// profile
	router.GET("/2/timeline/profile/{:userID}.json", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.FetchTweetsByUserIDLegacy(c.Param("userID"), scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetUserTweets via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user tweets"})
	})

	// conversation
	router.GET("/2/timeline/conversation/{:id}.json", func(c *gin.Context) {
		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetTweet1(scraper, c.Param("id"))
			if err != nil {
				logrus.Errorf("failed to GetTweet1 via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch conversation"})
	})

	// GetTweet2 - TweetDetail
	router.GET("/i/api/graphql/VWFGPVAGkZMGRKGe3GFFnA/TweetDetail", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetTweet2(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetTweet2 via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweet detail"})
	})

	// TweetResultByRestId
	router.GET("/i/api/graphql/xBtHv5-Xsk268T5ng_OGNg/TweetResultByRestId", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetTweetResultByRestId(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetTweetResultByRestId via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweet result by restId"})
	})

	// HomeLatestTimeline
	router.GET("/i/api/graphql/9EwYy8pLBOSFlEoSP2STiQ/HomeLatestTimeline", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetHomeLatestTimeline(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetHomeLatestTimeline via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweet result by restId"})
	})

	// HomeTimeline
	router.GET("/i/api/graphql/1u0Wlkw6Ru1NwBUD-pDiww/HomeTimeline", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetHomeTimeline(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetHomeLatestTimeline via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweet result by restId"})
	})

	// Bookmarks
	router.GET("/i/api/graphql/-IyJFt9_jS_9d_vS3NN-fA/Bookmarks", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetBookmarks(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetBookmarks via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookmarks"})
	})

	// Following
	router.GET("/i/api/graphql/g5P4cbXR4ta4oCeE7y2vLQ/Following", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetFollowing(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetFollowing via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch following"})
	})

	// Followers
	router.GET("/i/api/graphql/jwbfbSzn0FRL_AMZGsYDag/Followers", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetFollowers(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetFollowing via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch followers"})
	})

	// UserMedia
	router.GET("/i/api/graphql/2tLOJWwGuCTytDrGBg8VwQ/UserMedia", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetUserMedia(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetFollowing via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user media"})
	})

	// UserByScreenName
	router.GET("/graphql/Yka-W8dz7RaEuQNkroPkYw/UserByScreenName", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetUserByScreenName(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetFollowing via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user by username"})
	})

	// UserByRestId
	router.GET("/i/api/graphql/Qw77dDjp9xCpUY-AXwt-yQ/UserByRestId", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetUserByRestId(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetUserByRestId via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user by restId"})
	})

	// TweetDetail
	router.GET("/i/api/graphql/ldqoq5MmFHN1FhMGvzC9Jg/TweetDetail", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetTweetDetail(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetTweetDetail via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tweet detail"})
	})

	// FetchScheduledTweets
	router.GET("/i/api/graphql/ITtjAzvlZni2wWXwf295Qg/FetchScheduledTweets", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetFetchScheduledTweets(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetTweetDetail via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch scheduled tweets"})
	})

	// DeleteScheduledTweet
	router.POST("/i/api/graphql/CTOVqej0JBXAZSwkp1US0g/DeleteScheduledTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.DeleteScheduledTweet(scraper, payload)
			if err != nil {
				logrus.Errorf("failed to DeleteScheduledTweet via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete scheduled tweet"})
	})

	// CreateScheduledTweet
	router.POST("/i/api/graphql/LCVzRQGxOaGnOnYH01NQXg/CreateScheduledTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.CreateScheduledTweet(scraper, payload)
			if err != nil {
				logrus.Errorf("failed to CreateScheduledTweet via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create scheduled tweet"})
	})

	// SearchTimeline
	router.GET("/i/api/graphql/nK1dw4oV3k4w5TdtcAdSww/SearchTimeline", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.SearchTimeline(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to SearchTimeline via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search timeline"})
	})

	// AudioSpaceById
	router.GET("/i/api/graphql/d03OdorPdZ_sH9V3D1_yWQ/AudioSpaceById", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetAudioSpaceById(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetAudioSpaceById via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch audio space by id"})
	})

	// guide
	router.GET("/2/guide.json", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetGuide(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetGuide via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch guide"})
	})

	// CreateTweet
	router.POST("/i/api/graphql/oB-5XsHNAbjvARJEc8CZFw/CreateTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.CreateTweet(scraper, payload)
			if err != nil {
				logrus.Errorf("failed to CreateTweet via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tweet"})
	})

	// DeleteTweet
	router.POST("/i/api/graphql/VaenaVgh5q5ih7kvyVjgtg/DeleteTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.DeleteTweet(scraper, payload)
			if err != nil {
				logrus.Errorf("failed to DeleteTweet via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tweet"})
	})

	// CreateRetweet
	router.POST("/i/api/graphql/ojPdsZsimiJrUGLR1sjUtA/CreateRetweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.CreateRetweet(scraper, payload)
			if err != nil {
				logrus.Errorf("failed to CreateRetweet via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create retweet"})
	})

	// DeleteRetweet
	router.POST("/i/api/graphql/iQtK4dl5hBmXewYZuEOKVw/DeleteRetweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.DeleteRetweet(scraper, payload)
			if err != nil {
				logrus.Errorf("failed to DeleteRetweet via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete retweet"})
	})

	// FavoriteTweet
	router.POST("/i/api/graphql/lI07N6Otwv1PhnEgXILM7A/FavoriteTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.FavoriteTweet(scraper, payload)
			if err != nil {
				logrus.Errorf("failed to FavoriteTweet via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to favorite tweet"})
	})

	// UnfavoriteTweet
	router.POST("/i/api/graphql/ZYKSe-w7KEslx3JhSIk5LA/UnfavoriteTweet", func(c *gin.Context) {
		var payload map[string]interface{}
		err := c.BindJSON(&payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.UnfavoriteTweet(scraper, payload)
			if err != nil {
				logrus.Errorf("failed to UnfavoriteTweet via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unfavorite tweet"})
	})

	// GetRetweeters
	router.GET("/i/api/graphql/8019obfgnveiPiJuS2Rtow/Retweeters", func(c *gin.Context) {
		encodedQuery := c.Request.URL.Query().Encode()

		for index := 0; index < as.GetAccountsCount(); index++ {
			scraper, account, err := as.GetAuthenticatedScraper()
			if err != nil {
				logrus.Errorf("failed to get scraper of %s, %v", account.Username, err)
				continue
			}
			timeline, err := ts.GetRetweeters(scraper, encodedQuery)
			if err != nil {
				logrus.Errorf("failed to GetRetweeters via %s, %v", account.Username, err)
				as.HandleRateLimit(err, account)
				continue
			}
			c.JSON(http.StatusOK, timeline)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch retweeters"})
	})
}
