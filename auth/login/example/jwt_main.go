package main

import (
	"net/http"

	"github.com/your-username/your-project-name/auth"
)

/*
When a user accesses the /auth endpoint, the Authenticate function will be called, which extract the JWT from the request header, parse it and validate it using the secret key that you provide, if the token is not valid it will return a error message.
It's important to replace the placeholder value in the example above with the actual value of your secret key.



*/

func main() {
	jwtAuth := &auth.JWTAuth{Secret: []byte("your_secret_key")}

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		jwtAuth.Authenticate(w, r)

	})
	http.ListenAndServe(":8080", nil)

}
