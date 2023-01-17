package main

import (
	"fmt"
	"log"

	"github.com/SlyMarbo/rss"
)

/*
This example uses the Fetch function from the "github.com/SlyMarbo/rss" package to retrieve the ATOM feed from a URL. The returned feed struct contains the parsed data, with the Title and Items fields populated. The Items field is a slice of Item structs, which contain the entry title and link. Finally, it prints the title of the feed and each entry title and link in the feed.

Note that, the log package is used to log the error, which will give more information like file, line, timestamp and level of error.

You can use other functions like FetchByHTTPClient and FetchByClient to fetch the feed by providing a http client instance or FetchByBytes to fetch feed from already downloaded bytes.
*/
func main() {
	feed, err := rss.Fetch("https://example.com/feed.atom")
	if err != nil {
		log.Fatalf("Error fetching ATOM feed: %v", err)

	}

	fmt.Println("Title:", feed.Title)
	fmt.Println("Entries:")
	for _, entry := range feed.Items {
		fmt.Println("  Title:", entry.Title)
		fmt.Println("  Link:", entry.Link)

	}

}
