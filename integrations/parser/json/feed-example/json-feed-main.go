package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yourusername/jsonparser"
)

/*
the main function uses the http.Get function to get the feed data from a JSON file. It reads the response body using ioutil.ReadAll and parses the JSON data using the jsonparser.ParseJSON function.
Once the JSON data is parsed, it is stored in the variable feedData which is of type interface{} so that it can accept any type of json data.
It then displays the feed data using the fmt.Println function
*/
func main() {
	// Get the feed data
	resp, err := http.Get("https://www.example.com/feed.json")
	if err != nil {
		fmt.Println("Error getting feed:", err)
		return

	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return

	}

	// Parse the JSON data
	var feedData interface{}
	err = jsonparser.ParseJSON(bytes.NewReader(body), &feedData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return

	}

	// Display the feed data
	fmt.Println("Feed data:", feedData)

}
