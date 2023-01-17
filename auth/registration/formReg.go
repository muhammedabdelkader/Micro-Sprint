package users

import (
    "fmt"
        "net/http"
            "encoding/json"
            
       )

       // User represents a registered user
type User struct {
    ID       int    `json:"id"`
        Username string `json:"username"`
            Email    string `json:"email"`
                Password string `json:"password"`
                
}

// Users is a map that stores all registered users
var Users = map[int]User{}

// RegisterUser is a function that allows users to register
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user User
        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
                http.Error(w, "Invalid request body", http.StatusBadRequest)
                        return
                            
        }

    if user.Username == "" || user.Email == "" || user.Password == "" {
            http.Error(w, "Invalid request body, username, email and password
            fields are required", http.StatusBadRequest)
                    return
                        
    }

        // Assign an ID to the new user
            maxID := 0
                       for id := range Users {
                           if id > maxID {
                                       maxID = id
                                               
                           }
                               
                       }
                           user.ID = maxID + 1

                               // Store the new user
                                   Users[user.ID] = user

                                       // Return the new user
                                           json.NewEncoder(w).Encode(user)
                                               fmt.Fprintln(w, "User registered
                                               successfully!")
                                               
}

