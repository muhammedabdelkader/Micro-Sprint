package main

import (
	"net/http"

	"github.com/your-username/your-project-name/auth"
)

/*
When a user accesses the /auth endpoint, the GithubAuth function will be called. This function will redirect the user to the Github login page, where they can sign in with their Github credentials. Once the user is logged in, the callback function will be executed and you can use the goth.User object to access the user's information.

You need to replace the client-id and client-secret with the actual values obtained from the Github developer portal.
*/
func main() {
	http.HandleFunc("/auth", auth.GithubAuth)
	http.ListenAndServe(":8080", nil)

}
