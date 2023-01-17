package main

import (
	"net/http"

	"github.com/your-username/your-project-name/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/oidc"
)

/*
When a user accesses the /login endpoint, the application will redirect them to the OpenID Connect provider's login page, where they can sign in with their credentials.
Once the user is logged in
*/
func main() {
	// Create the OAuth2 config
	oidcConfig := &oidc.Config{
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
		RedirectURL:  "http://localhost:8080/callback",
		ProviderURL:  "https://your_provider_url",
		Scopes:       []string{"openid", "profile", "email"},
	}
	oidcAuth := &auth.OIDCAuth{OAuthConfig: oauth2.Config{
		ClientID:     oidcConfig.ClientID,
		ClientSecret: oidcConfig.ClientSecret,
		RedirectURL:  oidcConfig.RedirectURL,
		Endpoint:     oidc.Endpoint,
		Scopes:       oidcConfig.Scopes,
	},
	}

	// Redirect the user to the provider's login page
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, oidcAuth.OAuthConfig.AuthCodeURL(""), http.StatusSeeOther)

	})

	// Handle the callback from the provider
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		oidcAuth.Authenticate(w, r)

	})

	http.ListenAndServe(":8080", nil)

}
