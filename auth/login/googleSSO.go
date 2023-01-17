package auth

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

/*
You can use this module by importing it in your application and calling the GoogleAuth function with an http.ResponseWriter and http.Request as arguments. This will redirect the user to the Google login page, where they can sign in with their Google credentials. Once the user is logged in, the callback function will be executed and you can use the goth.User object to access the user's information.

You need to replace the client-id and client-secret with the actual values obtained from the Google developer console.
*/
var (
	googleProvider = google.New("client-id", "client-secret", "http://localhost:8080/auth/google/callback")
)

func init() {
	goth.UseProviders(googleProvider)

}

func GoogleAuth(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/auth/google/callback" {
		user, err := goth.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return

		}

		// Perform any additional actions here, such as storing
		// the user's information in a database
		fmt.Fprintln(w, "Welcome,", user.Name)
		return

	}

	goth.BeginAuthHandler(w, r)

}
