package auth

import (
	"errors"
)

/*
You can use this module by importing it in your application and calling the Authenticate function with a username and password.
It returns a boolean value indicating whether the authentication was successful and an error message in case of failure.
*/
// User represents a user with a username and password
type User struct {
	Username string
	Password string
}

// Users is a map of registered users
var Users = map[string]User{
	"user1": {"user1", "pass1"},
	"user2": {"user2", "pass2"},
}

// Authenticate checks if a username and password match a registered user
func Authenticate(username, password string) (bool, error) {
	user, ok := Users[username]
	if !ok {
		return false, errors.New("invalid username")

	}

	if user.Password != password {
		return false, errors.New("invalid password")

	}

	return true, nil

}
