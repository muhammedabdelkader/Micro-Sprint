package main

import (
	"fmt"
	"net/http"
)
/*
In this code, the user role is checked against a hard-coded value and 
if it matches, the user is granted access to the sensitive function of 
deleting all users in the system. An attacker could manipulate the user 
role value by passing a malicious input, such as "admin" in order to gain unauthorized access to the system.
*/
func main() {
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		userRole := r.FormValue("role")
		if userRole == "admin" {
			deleteAllUsers()
			fmt.Fprintf(w, "All users have been deleted.")
		} else {
			fmt.Fprintf(w, "Access denied.")
		}
	})
	http.ListenAndServe(":8080", nil)
}

func deleteAllUsers() {
	// code to delete all users from the database
}
