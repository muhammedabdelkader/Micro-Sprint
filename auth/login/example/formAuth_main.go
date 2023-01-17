package main

import (
	"net/http"

	"github.com/your-username/your-project-name/auth"
)

/*
When a user accesses the /auth endpoint, the UsernameAuth function will be called. This function will extract the username and password from the form values using r.FormValue("username") and r.FormValue("password") respectively, then it will compare them against a predefined list of usernames and passwords
*/
func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		auth.UsernameAuth(w, r)

	})
	http.ListenAndServe(":8080", nil)

}
