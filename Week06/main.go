package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
)

func main() {
	// OAuth1 config
	config := oauth1.NewConfig("IcdJMJ0rD6oCzpgUSIbk2eGHG", "QfuFGCtnyzgsLJgreOJ6mPyI5jrhguCxQcVnRYrnMGQuIoIz1Y")
	token := oauth1.NewToken("2170385593-8lQImNPXJAmlRpTl56AyUYwfcv94C4tunNSCjUw", "0Jv3lnOlg7pma9ptUu2DKKnLc4CBFPd4uyNmPkwkTK1YM")
	httpClient := config.Client(oauth1.NoContext, token)

	// Set tweet content
	tweetURL := "https://api.twitter.com/1.1/statuses/update.json"
	data := url.Values{}
	data.Set("status", "Hello from Twitter API!")

	// Make the POST request
	resp, err := httpClient.PostForm(tweetURL, data)
	if err != nil {
		log.Fatalf("Error posting tweet: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Tweet posted successfully!")
	} else {
		fmt.Printf("Failed to post tweet, status code: %d\n", resp.StatusCode)
	}
}
