package auth

import (
	"fmt"
	"net/http"

	"github.com/vjeantet/ldapserver"
)

/*
You can use this module by importing it in your application and calling the ActiveDirectoryAuth function with an http.ResponseWriter and http.Request as arguments. This function will extract the username and password from the form values and authenticate against the active directory server using the ldapserver package. Once the user is logged in, the callback function will be executed
*/
// ActiveDirectoryAuth performs authentication using SSO with Active Directory
func ActiveDirectoryAuth(w http.ResponseWriter, r *http.Request) {
	ldapServer := ldapserver.NewServer()
	defer ldapServer.Stop()

	authenticated, username, err := ldapServer.Authenticate(r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		fmt.Fprintln(w, err)
		return

	}

	if !authenticated {
		http.Redirect(w, r, "/login", http.StatusFound)
		return

	}

	// Perform any additional actions here, such as storing the
	// user's information in a database
	fmt.Fprintln(w, "Welcome,", username)

}
