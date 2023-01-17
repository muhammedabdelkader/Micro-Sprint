package fbbot

import (
	"fmt"
	"net/url"

	"github.com/huandu/facebook"
)

/*
the package defines a struct FacebookClient that holds the facebook-go-sdk client and it also defines two methods GetAccessToken
*/
// FacebookClient is a struct that holds the facebook client
type FacebookClient struct {
	client *facebook.Facebook
}

// NewFacebookClient creates and returns a new FacebookClient
func NewFacebookClient(appID, appSecret string) (*FacebookClient, error) {
	client := facebook.New(appID, appSecret)
	return &FacebookClient{client: client}, nil

}

// GetAccessToken gets an access token
func (f *FacebookClient) GetAccessToken() error {
	res, err := f.client.Get("/oauth/access_token", url.Values{
		"client_id":     {f.client.AppID},
		"client_secret": {f.client.AppSecret},
		"grant_type":    {"client_credentials"},
	})
	if err != nil {
		return fmt.Errorf("Error getting access token: %s", err)

	}
	f.client.SetAccessToken(res["access_token"].(string))
	return nil

}

// PostMessage posts a message on facebook
func (f *FacebookClient) PostMessage(message string) error {
	_, err := f.client.Post("/me/feed", url.Values{
		"message": {message},
	})
	if err != nil {
		return fmt.Errorf("Error posting message: %s", err)

	}
	return nil

}

// DeletePost deletes a post on Facebook
/*
This function takes a postID as a parameter, it uses the f.client.Delete(postID) function to delete the post on Facebook.
It returns error if it fails otherwise it will return nil
*/
func (f *FacebookClient) DeletePost(postID string) error {
	_, err := f.client.Delete(postID)
	if err != nil {
		return fmt.Errorf("Error deleting post: %s", err)

	}
	return nil

}

// ChangePostPrivacy changes the privacy of a post on Facebook
/*
This function takes a postID and a privacy as parameters, it uses the f.client.Post(postID, url.Values{"privacy": {privacy}}) function to change the privacy of the post on Facebook.
It returns error if it fails otherwise it will return nil
The privacy parameter can take values such as EVERYONE, FRIENDS or CUSTOM and you can find more options of privacy on the Facebook documentation
*/
func (f *FacebookClient) ChangePostPrivacy(postID string, privacy string) error {
	_, err := f.client.Post(postID, url.Values{
		"privacy": {privacy},
	})
	if err != nil {
		return fmt.Errorf("Error changing privacy of post: %s", err)

	}
	return nil

}
