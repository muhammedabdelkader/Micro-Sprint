package main

import (
	"fmt"
	"log"

	"github.com/mmcdole/gofeed"
)

/*
This example uses the gofeed.NewParser() function to create a new parser, and the ParseURL function to parse the ATOM feed from a given URL. The returned feed struct contains the parsed data, with the Title and Items fields populated. The Items field is a slice of Item structs, which contain the entry title and link. Finally, it prints the title of the feed and each entry title and link in the feed.

Note that, the log package is used to log the error, which will give more information like file, line, timestamp and level of error.

You can also use Parse to parse feed from bytes or ParseString to parse feed from string.

It also supports parsing of RSS and RDF feed types and can be used to parse different feed types as well.
*/
func main() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://example.com/feed.atom")
	if err != nil {
		log.Fatalf("Error parsing ATOM feed: %v", err)

	}

	fmt.Println("Title:", feed.Title)
	fmt.Println("Entries:")
	for _, entry := range feed.Items {
		fmt.Println("  Title:", entry.Title)
		fmt.Println("  Link:", entry.Link)

	}

}
