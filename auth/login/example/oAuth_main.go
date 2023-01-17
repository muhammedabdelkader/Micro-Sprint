package main

import (
	"net/http"

	"github.com/your-username/your-project-name/auth"
)

/*
When a user accesses the /auth endpoint, the OAuthAuth function will be called, which creates an OAuthConfig struct, using the provided ClientID, ClientSecret, TokenURL and Scopes values, then it uses this config to get an access token from the OAuth2 provider.

It's important to note that you should replace the placeholder values in the example above with the actual values that correspond to your OAuth2 provider (e.g. client_id and client_secret). Also, you should replace the token URL with the URL provided by your OAuth2 provider.



*/

func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		config := &auth.OAuthConfig{
			ClientID:     "your_client_id",
			ClientSecret: "your_client_secret",
			TokenURL:     "https://your-provider.com/token",
			Scopes:       []string{"scope1", "scope2"},
		}
		auth.OAuthAuth(w, r, config)

	})
	http.ListenAndServe(":8080", nil)

}
