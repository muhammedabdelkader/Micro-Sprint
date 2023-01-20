package main

import (
	"fmt"
	"net/http"
)

/*
In this code, the user input is not properly validated or sanitized 
before being used in the HTML output. 
An attacker could provide a malicious input, 
such as <script>alert('XSS')</script> in order to execute arbitrary 
JavaScript code in the victim's browser.
*/

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		userInput := r.FormValue("name")
		html := fmt.Sprintf(`<html><body><h1>Hello, %s!</h1></body></html>`, userInput)
		fmt.Fprintf(w, html)
	})
	http.ListenAndServe(":8080", nil)
}
