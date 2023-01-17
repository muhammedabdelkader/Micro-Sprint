package auth

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

/*

 */
// OAuthConfig is a struct that contains the OAuth2 configuration
type OAuthConfig struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
	Scopes       []string
}

// OAuthAuth performs authentication using OAuth2 with a generic provider
func OAuthAuth(w http.ResponseWriter, r *http.Request, config *OAuthConfig) {
	// Create the OAuth2 config
	oauthConfig := &clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		TokenURL:     config.TokenURL,
		Scopes:       config.Scopes,
	}

	// Get the token
	token, err := oauthConfig.Token(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	// Perform any additional actions here, such as storing the
	// user's information in a database
	fmt.Fprintln(w, "Access Token:", token.AccessToken)

}
