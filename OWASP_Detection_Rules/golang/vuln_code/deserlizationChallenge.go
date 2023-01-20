package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func updatePerson(input []byte) {
	var person Person
	json.Unmarshal(input, &person)
	fmt.Println("Updating person:", person)

	// do some critical update, for example update the person in the database
	// ...
}

func main() {
	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		input, _ := ioutil.ReadAll(r.Body)
		updatePerson(input)
		w.Write([]byte("Success"))
	})
	http.ListenAndServe(":8080", nil)
}


/*
Solution: 
payload that contains a Go expression that gets executed as soon as it gets deserialized,  
This payload would be deserialized into the Person struct and the reflect.ValueOf field would execute the command os.Remove("/etc/passwd"), deleting the password file on the target system.

To prevent this type of attack, it's important to validate the input before deserializing it and also use a library that has been specifically designed to mitigate deserialization attacks, such as the go-unsafe-json package. Additionally, it is recommended to use a strict JSON decoding mode and/or use a JSON decoder that is configured to only decode specific fields.

{
	"Name":"John Smith",
	"Age":30,
	"Address":"http://attacker.com",
	"reflect.ValueOf": "`os.Remove(\"/etc/passwd\")`"
}

*/