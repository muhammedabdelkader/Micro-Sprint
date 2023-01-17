package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
In this version, I've added error handling to check for any errors that may occur when making the HTTP request, reading the response body, and unmarshalling the XML data. If an error occurs, it prints an error message and returns from the function, instead of continuing execution and potentially causing a panic.

Also, you can use log package instead of fmt to log the error, which will give more information like file, line, timestamp and level of error.
*/
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Title   string   `xml:"title"`
	Entry   []Entry  `xml:"entry"`
}

type Entry struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

func main() {
	resp, err := http.Get("https://example.com/feed.atom")
	if err != nil {
		fmt.Println("Error getting ATOM feed:", err)
		return

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading ATOM feed:", err)
		return

	}

	var feed Feed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		fmt.Println("Error parsing ATOM feed:", err)
		return

	}

	fmt.Println("Title:", feed.Title)
	fmt.Println("Entries:")
	for _, entry := range feed.Entry {
		fmt.Println("  Title:", entry.Title)
		fmt.Println("  Link:", entry.Link)

	}

}
