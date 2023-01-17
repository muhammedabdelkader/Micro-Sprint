package news

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
The `NewsClient` struct is used to define an API key, tokens, and base URL, as well as an http client to make requests to the news API. The `NewClient` function is used to initialize the `NewsClient` struct with an API key and any necessary tokens.

The `Query` method takes an endpoint and a map of parameters as input and makes a GET request to the news API using the `BaseURL`, `APIKey`, and `Tokens` fields of the `NewsClient` struct. It then encodes the query parameters into the URL and returns the response body as a byte slice.

The `GetHotNews()` and `GetLatestNewsInCategory(category string)` functions use the `Query` method to make requests to the news API, passing the appropriate endpoint and parameters, and then unmarshal the JSON response into the `News` struct.
*/
type News struct {
	Articles []Article `json:"articles"`
}

type Article struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type NewsClient struct {
	APIKey  string
	Tokens  map[string]string
	BaseURL string
	Client  *http.Client
}

func NewClient(apiKey string) *NewsClient {
	tokens := make(map[string]string)
	// Add any tokens you need to the map
	return &NewsClient{
		APIKey:  apiKey,
		Tokens:  tokens,
		BaseURL: "https://newsapi.org/v2/",
		Client:  &http.Client{},
	}

}

func (nc *NewsClient) Query(endpoint string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", nc.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err

	}

	q := req.URL.Query()
	q.Add("apiKey", nc.APIKey)
	for k, v := range nc.Tokens {
		q.Add(k, v)

	}
	for k, v := range params {
		q.Add(k, v)

	}
	req.URL.RawQuery = q.Encode()

	resp, err := nc.Client.Do(req)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}

	return body, nil

}

// GetHotNews returns a list of hot news articles
func (nc *NewsClient) GetHotNews() (News, error) {
	var hotNews News
	body, err := nc.Query("top-headlines", map[string]string{
		"country": "us",
		"sortBy":  "popularity",
	})
	if err != nil {
		return hotNews, err

	}

	err = json.Unmarshal(body, &hotNews)
	if err != nil {
		return hotNews, err

	}

	return hotNews, nil

}

// GetLatestNewsInCategory returns a list of latest news articles in a specific
// category
func (nc *NewsClient) GetLatestNewsInCategory(category string) (News, error) {
	var latestNews News
	body, err := nc.Query("top-headlines", map[string]string{
		"country":  "us",
		"category": category,
	})
	if err != nil {
		return latestNews, err

	}

	err = json.Unmarshal(body, &latestNews)
	if err != nil {
		return latest, nil

	}
	return latestNews, nil

}
