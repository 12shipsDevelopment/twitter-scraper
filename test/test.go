package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	twitterscraper "github.com/12shipsDevelopment/twitter-scraper"
)

func main() {
	scraper := twitterscraper.New()
	scraper.WithClientTimeout(time.Second * 100)
	// scraper.SetAuthToken(twitterscraper.AuthToken{Token: "56d59f2b5af98a334de8f4cff738cd2335955b2b", CSRFToken: "549bf61a173d835422bef4d79d58d66ccb81611b8fb915c17905a69e77528e4e0089caeb6119639ac643f2c4f8b4bca09bd1107a56342b043c093d59f43f12eb30bb18b70fabbcb4ff2e80bd8412cf5d"})

	// After setting Cookies or AuthToken you have to execute IsLoggedIn method.
	// Without it, scraper wouldn't be able to make requests that requires authentication
	var cookies []*http.Cookie
	f, _ := os.Open("cookies.json")

	json.NewDecoder(f).Decode(&cookies)
	fmt.Println(
		"Cookies: ", cookies,
	)
	err := scraper.SetProxy("http://localhost:7890")
	if err != nil {
		panic("Failed to set proxy, " + err.Error())
	}

	scraper.SetCookies(cookies)
	if !scraper.IsLoggedIn() {
		panic("Invalid cookies")
	}

	// t, err := scraper.GetTweet("1921263370503663770")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(t)
	for tweet := range scraper.SearchTweetsForSubnet13(context.Background(), "x", 50) {
		if tweet.Error != nil {
			panic(tweet.Error)
		}
		b, err := json.MarshalIndent(tweet, "", " ")
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Print(string(b))
	}
}
