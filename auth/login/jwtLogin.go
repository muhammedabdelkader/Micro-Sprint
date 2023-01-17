package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

/*
You can use this module by importing it in your application and creating an instance of the JWTAuth struct, and then call the Authenticate function with an http.ResponseWriter and http.Request as arguments.
This function will extract the JWT from the request header, parse it and validate it using the secret key that

*/
// JWTAuth is a struct that contains the JWT configuration
type JWTAuth struct {
	Secret []byte
}

// Authenticate performs authentication using JWT
func (j *JWTAuth) Authenticate(w http.ResponseWriter, r *http.Request) {
	// Extract the JWT from the request
	tokenString := r.Header.Get("Authorization")

	// Parse the JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])

		}
		return j.Secret, nil

	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return

	}

	if !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return

	}

	// Perform any additional actions here, such as storing the
	// user's information in a database
	fmt.Fprintln(w, "Authenticated")

}
