package main

import (
	"net/http"

	"github.com/your-username/your-project-name/auth"
)

/*
When a user accesses the /auth endpoint, the ActiveDirectoryAuth function will be called. This function will extract the username and password from the form values using r.FormValue("username") and r.FormValue("password") respectively, then it will authenticate against the active directory server using the ldapserver.Authenticate function, passing the username and password as arguments.
*/
func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		auth.ActiveDirectoryAuth(w, r)

	})
	http.ListenAndServe(":8080", nil)

}
