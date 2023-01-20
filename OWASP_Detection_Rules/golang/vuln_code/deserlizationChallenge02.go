package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Name    string
	Address string
	Command string
}

func handleRequest(req Request) {
	fmt.Println("Received request from", req.Name)
	fmt.Println("Updating address to", req.Address)

	// execute command
	result, err := exec.Command(req.Command).Output()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(string(result))
}

func main() {
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req Request
		err := decoder.Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}
		handleRequest(req)
		w.Write([]byte("Success"))
	})
	http.ListenAndServe(":8080", nil)
}

/*
Solution: 
This application listens to HTTP requests on port 8080, when it receives a request to the "/update" endpoint, it deserializes the JSON payload into the Request struct, and the handleRequest function updates the address and execute the command.

An attacker could exploit this vulnerability by sending a malicious payload that contains a command that gets executed as soon as it gets deserialized, 

This payload would be deserialized into the Request struct and the handleRequest function would execute the command rm -rf /, which would delete the entire file system on the target system

{
	"Name":"John Smith",
	"Address":"http://attacker.com",
	"Command":"rm -rf /"
}


*/