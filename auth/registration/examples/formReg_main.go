package main

import (
	"fmt"
	"net/http"
)

/*
You should call this function to register a user by sending a POST request with JSON data containing the fields "username", "email", and "password" to the "/register" endpoint.

Also you can test it by using tool like postman or curl by sending request with json body to endpoint /register

You can find more information on how to set up an HTTP server and handle requests in the Go documentation for the "net/http" package.
*/
func main() {
	http.HandleFunc("/register", RegisterUser)
	fmt.Println("Starting server on port 3000...")
	http.ListenAndServe(":3000", nil)

}
