package auth

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/oidc"
)

/*
You can use this function by importing it in your application and creating an instance of the OIDCAuth struct and configuring it with the appropriate values, such as client ID, client secret, redirect URL, and OpenID Connect provider's URL.
Then, you can use the Authenticate function in your HTTP handler to perform authentication.
In addition, different OpenID providers have different ways of providing the client id and secret, redirect url and provider url, you should check the provider's documentation to get the correct values.

*/
// OIDCAuth is a struct that contains the OpenID Connect configuration
type OIDCAuth struct {
	OAuthConfig *oauth2.Config
}

// Authenticate performs authentication using OpenID Connect
func (o *OIDCAuth) Authenticate(w http.ResponseWriter, r *http.Request) {
	// Get the OAuth2 token from the request
	token, err := o.OAuthConfig.Exchange(context.TODO(), r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	// Verify the token and extract the claims
	verifier := o.OAuthConfig.Provider.(*oidc.Provider).Verifier(oidc.Config{ClientID: o.OAuthConfig.ClientID})
	claims, err := verifier.Verify(context.TODO(), token.Extra("id_token").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	// Perform any additional actions here, such as
	// storing the user's information in a database
	fmt.Fprintln(w, "Authenticated")
	fmt.Println("User ID:", claims.Subject)

}
