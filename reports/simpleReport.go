package main

import (
	"fmt"
	"net/http"
)

/*
In this example, the function submitSecurityReport is handling the form submission and it will check the request method if it is not post it will return error. Then it will parse the form and validate the fields, if any fields are empty it will return error. Then it will create a new security report struct and appends it to SecurityReports slice. The slice is used to store all the security reports.

You can also use r.PostForm to parse the form and validate the fields, it will return a map of form fields and values.

You can add more security checks like validate the email field or encrypting the password before storing it in the struct.

You can test this by sending a post request to the endpoint "/submit-security-report" with form fields "title", "description", "reporter", and "email".

*/
// SecurityReport represents a security report
type SecurityReport struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Reporter    string `json:"reporter"`
	Email       string `json:"email"`
}

// SecurityReports is a slice that stores all security reports
var SecurityReports []SecurityReport

/*
	func submitSecurityReport(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return

		}

		// Parse the form
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return

		}

		// Validate the form fields
		title := r.FormValue("title")
		description := r.FormValue("description")
		reporter := r.FormValue("reporter")
		email := r.FormValue("email")

		if title == "" || description == "" || reporter == "" || email == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return

		}

		// Create a new security report
		securityReport := SecurityReport{
			Title:       title,
			Description: description,
			Reporter:    reporter,
			Email:       email,
		}

		// Append the security report to the
		// slice
		SecurityReports = append(SecurityReports, securityReport)

		fmt.Fprintln(w, "Security report submitted successfully!")

}
*/

/*
the function first verifies that the user is logged in by checking the "loggedin" value in the session. The session is created using a package like gorilla/sessions and it uses a session store like memory, file system, or redis. Once the session is verified to be true, the request will be parsed and validated and the security report will be created and appended to the slice.

You can also use a JWT token to verify the user is logged in. Once the user is logged in, you will generate a token and store it in the client side and then check the token in every request to verify the user is logged in.
*/
func submitSecurityReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return

	}
	// Verify that the user is logged in
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return

	}
	if loggedIn, ok := session.Values["loggedin"].(bool); !ok || !loggedIn {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return

	}
	// Parse the form
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return

	}
	// Validate the form fields
	title := r.FormValue("title")
	description := r.FormValue("description")
	reporter := r.FormValue("reporter")
	email := r.FormValue("email")
	if title == "" || description == "" || reporter == "" || email == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return

	}
	// Create a new security report
	securityReport := SecurityReport{
		Title:       title,
		Description: description,
		Reporter:    reporter,
		Email:       email,
	}
	// Append the security report to
	// the slice
	SecurityReports = append(SecurityReports, securityReport)
	fmt.Fprintln(w, "Security report submitted successfully!")

}

func main() {
	http.HandleFunc("/submit-security-report", submitSecurityReport)
	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", nil)

}
