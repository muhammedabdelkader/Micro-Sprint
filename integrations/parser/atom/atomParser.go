package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
This example uses the http package to retrieve the ATOM feed from a URL, then uses the xml package to parse the feed into a struct. The structs are defined using the xml tags to specify how the data should be mapped to the struct fields. The Title and Entry fields of the Feed struct, and the Title and Link fields of the Entry struct are populated. Finally, it prints the title of the feed and each entry title and link in the feed.

Note: This code snippet is just an example, and it does not handle error properly, it is just for demonstration purpose.

Also, you can use other packages available like github.com/mmcdole/gofeed or github.com/SlyMarbo/rss to parse ATOM feed which are very popular and easy to use
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
	resp, _ := http.Get("https://example.com/feed.atom")
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var feed Feed
	xml.Unmarshal(body, &feed)

	fmt.Println("Title:", feed.Title)
	fmt.Println("Entries:")
	for _, entry := range feed.Entry {
		fmt.Println("  Title:", entry.Title)
		fmt.Println("  Link:", entry.Link)

	}

}
