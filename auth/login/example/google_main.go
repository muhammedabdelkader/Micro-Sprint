package main

import (
	"net/http"

	"github.com/your-username/your-project-name/auth"
)

/*
When a user accesses the /auth endpoint, the GoogleAuth function will be called. This function will redirect the user to the Google login page, where they can sign in with their Google credentials. Once the user is logged in, the callback function will be executed and you can use the goth.User object to access the user's information.

You need to replace the client-id and client-secret with the actual values obtained from the Google developer console.
*/
func main() {
	http.HandleFunc("/auth", auth.GoogleAuth)
	http.ListenAndServe(":8080", nil)

}
