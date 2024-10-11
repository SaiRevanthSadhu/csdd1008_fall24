package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"io"

	"github.com/dghubble/oauth1"
)

// Function to post a new tweet using dghubble/oauth1
func postTweet(apiURL, tweetContent string, config *oauth1.Config, token *oauth1.Token) {
	// Create an HTTP client with OAuth 1.0a tokens
	httpClient := config.Client(oauth1.NoContext, token)

	// Prepare the data for the request
	data := url.Values{}
	data.Set("status", tweetContent)

	// Create a POST request to post a tweet
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Error creating request: %v\n", err)
	}

	// Set content type to application/x-www-form-urlencoded
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v\n", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, _ := io.ReadAll(resp.Body)

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to post tweet, status code: %d\n", resp.StatusCode)
		log.Printf("Response: %s\n", string(body))
		return
	}

	fmt.Println("Tweet posted successfully!")
	fmt.Printf("Response: %s\n", string(body)) // Print the response (you can parse the tweet ID from here)
}

// Function to delete a tweet using dghubble/oauth1
func deleteTweet(apiURL, tweetID string, config *oauth1.Config, token *oauth1.Token) {
	// Create an HTTP client with OAuth 1.0a tokens
	httpClient := config.Client(oauth1.NoContext, token)

	// Construct the DELETE request
	apiEndpoint := fmt.Sprintf("%s/%s.json", apiURL, tweetID)
	req, err := http.NewRequest("POST", apiEndpoint, nil) // Twitter uses POST for delete
	if err != nil {
		log.Fatalf("Error creating request: %v\n", err)
	}

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v\n", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, _ := io.ReadAll(resp.Body)

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to delete tweet, status code: %d\n", resp.StatusCode)
		log.Printf("Response: %s\n", string(body))
		return
	}

	fmt.Println("Tweet deleted successfully!")
	fmt.Printf("Response: %s\n", string(body))
}

func main() {
	// Twitter API credentials
	consumerKey := os.Getenv("IcdJMJ0rD6oCzpgUSIbk2eGHG")
	consumerSecret := os.Getenv("QfuFGCtnyzgsLJgreOJ6mPyI5jrhguCxQcVnRYrnMGQuIoIz1Y")
	accessToken := os.Getenv("2170385593-8lQImNPXJAmlRpTl56AyUYwfcv94C4tunNSCjUw")
	accessSecret := os.Getenv("0Jv3lnOlg7pma9ptUu2DKKnLc4CBFPd4uyNmPkwkTK1YM")

	// OAuth1 configuration
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	// Post a tweet
	apiPostURL := "https://api.twitter.com/1.1/statuses/update.json"
	tweetContent := "Hello from the Twitter API using Go and oauth1!"
	postTweet(apiPostURL, tweetContent, config, token)

	// Example tweet ID for deletion (you can replace it with an actual tweet ID you get after posting)
	tweetID := "YOUR_TWEET_ID" // You need to replace this with the actual tweet ID
	apiDeleteURL := "https://api.twitter.com/1.1/statuses/destroy"
	deleteTweet(apiDeleteURL, tweetID, config, token)
}
