package twitterbot

import (
	"fmt"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

/*
the package defines a struct TwitterClient that holds the anaconda Twitter API client and it also defines two methods PostTweet and GetTimeline that can be used to interact with Twitter.

NewTwitterClient function creates and returns a new TwitterClient, it sets the consumer key and secret and the access token and secret of your account and creates a new twitter api client using the anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")) function.

PostTweet function sends a tweet using the message passed as a parameter and GetTimeline function gets the recent tweets from a user's timeline using the username passed as a parameter.


*/
// TwitterClient is a struct that holds the anaconda Twitter API client
type TwitterClient struct {
	api *anaconda.TwitterApi
}

// NewTwitterClient creates and returns a new TwitterClient
func NewTwitterClient() (*TwitterClient, error) {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return &TwitterClient{api: api}, nil

}

// PostTweet sends a tweet
func (t *TwitterClient) PostTweet(message string) error {
	_, err := t.api.PostTweet(message, nil)
	if err != nil {
		return fmt.Errorf("Error sending tweet: %s", err)

	}
	return nil

}

// GetTimeline gets the recent tweets from a user's timeline
func (t *TwitterClient) GetTimeline(username string) ([]anaconda.Tweet, error) {
	timeline, err := t.api.GetUserTimeline(username, nil)
	if err != nil {
		return nil, fmt.Errorf("Error getting timeline: %s", err)

	}
	return timeline, nil

}

// SendDM sends a direct message to a user
/*
This function takes a recipientUsername and a message as parameters, it creates a url values and set the recipientUsername and then sends the message as a direct message to the user using the t.api.PostDMToScreenName(message, v) function.
It returns error if it fails otherwise it will return nil.

*/
func (t *TwitterClient) SendDM(recipientUsername string, message string) error {
	v := url.Values{}
	v.Set("screen_name", recipientUsername)
	_, err := t.api.PostDMToScreenName(message, v)
	if err != nil {
		return fmt.Errorf("Error sending DM: %s", err)

	}
	return nil

}

// ListTweets returns a list of tweets that match a certain hashtag or from a
// certain user
/*
This function takes a searchTerm and a tweetType as parameters, it creates a url values and set the count of tweets to 100, it then gets tweets according to the tweetType passed as a parameter, whether it is a hashtag or a user tweets by calling the t.api.GetSearch(searchTerm, v) and t.api.GetUserTimeline(searchTerm, v) respectively.
It returns error if it fails otherwise it will return the tweets.
*/
func (t *TwitterClient) ListTweets(searchTerm string, tweetType string) ([]anaconda.Tweet, error) {
	v := url.Values{}
	v.Set("count", "100")
	var tweets []anaconda.Tweet
	var err error
	switch tweetType {
	case "hashtag":
		tweets, err = t.api.GetSearch(searchTerm, v)
	case "user":
		tweets, err = t.api.GetUserTimeline(searchTerm, v)
	default:
		return nil, fmt.Errorf("Invalid tweet type, should be either 'hashtag' or 'user'")

	}
	if err != nil {
		return nil, fmt.Errorf("Error getting tweets: %s", err)

	}
	return tweets, nil

}
