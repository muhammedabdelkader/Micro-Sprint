package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	ItemList    []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

/*
This module uses the encoding/xml package to parse the RSS feed and the net/http package to fetch it. It defines structs for the RSS, channel, and item elements, which are used to unmarshal the XML data into Go structs. The main function fetches the RSS feed, reads it, parses it, and prints the title, link, and description of each item in the feed. You can modify the structs and the logic of the main function to suit your needs
*/
func main() {
	// Replace with the URL of the RSS feed you want to parse
	url := "https://www.example.com/rss"

	// Fetch the RSS feed
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching feed:", err)
		return

	}
	defer resp.Body.Close()

	// Read the RSS feed
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading feed:", err)
		return

	}

	// Parse the RSS feed
	var rss RSS
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Println("Error parsing feed:", err)
		return

	}

	// Print the title, link, and
	// description of each item in
	// the RSS feed
	for _, item := range rss.Channel.ItemList {
		fmt.Println("Title:", item.Title)
		fmt.Println("Link:", item.Link)
		fmt.Println("Description:", item.Description)

	}

}
