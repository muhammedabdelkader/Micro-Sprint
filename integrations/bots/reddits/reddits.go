package reddit

import (
	"context"
	"log"
	"os"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

// Client struct to hold the reddit client
type Client struct {
	RedditClient *reddit.Bot
}

// NewClient returns a new reddit client
func NewClient(username, password, clientID, clientSecret, userAgent string) *Client {

	cfg := graw.Config{
		Subreddits: []string{"all"},
		UserAgent:  userAgent,
		Logger:     log.New(os.Stdout, "", log.LstdFlags),
	}

	harvest, err := reddit.NewBotHarvest(context.Background(), username, password, clientID, clientSecret, cfg)
	if err != nil {
		log.Fatal(err)

	}
	return &Client{harvest.Bot}

}

// ListSubreddit lists the subreddit by name
func (c *Client) ListSubreddit(subreddit string) ([]*reddit.Post, error) {
	posts, err := c.RedditClient.Listing(context.Background(), subreddit, "")
	if err != nil {
		return nil, err

	}
	return posts, nil

}

// SubmitPost submit a post to subreddit
func (c *Client) SubmitPost(subreddit, title, body string) (*reddit.Post, error) {
	post := reddit.Post{
		Title: title,
		Body:  body,
	}
	post, err := c.RedditClient.Submit(context.Background(), subreddit, post)
	if err != nil {
		return nil, err

	}
	return post, nil

}

// DeletePost deletes a post on subreddit
func (c *Client) DeletePost(postID string) error {
	err := c.RedditClient.Delete(context.Background(), postID)
	if err != nil {
		return err

	}
	return nil

}

// UpdatePost updates a post on subreddit
func (c *Client) UpdatePost(postID, newTitle, newText string) (*reddit.Post, error) {
	post, err := c.RedditClient.Edit(context.Background(), postID, newTitle, newText)
	if err != nil {
		return nil, err

	}
	return post, nil

}

// MoveUp a post in subreddit
func (c *Client) MoveUp(postID string) error {
	err := c.RedditClient.Upvote(context.Background(), postID)
	if err != nil {
		return err

	}
	return nil

}

// MoveDown a post in subreddit
func (c *Client) MoveDown(postID string) error {
	err := c.RedditClient.Downvote(context.Background(), postID)
	if err != nil {
		return err

	}
	return nil

}

// ListHotPosts lists hot posts on subreddit
func (c *Client) ListHotPosts(subreddit string) ([]*reddit.Post, error) {
	posts, err := c.RedditClient.Listing(context.Background(), subreddit, "hot")
	if err != nil {
		return nil, err

	}
	return posts, nil

}

// FetchLatestPostWithHighestLikes fetch latest post in subreddit sorted by
// highest likes
func (c *Client) FetchLatestPostWithHighestLikes(subreddit string) (*reddit.Post, error) {
	posts, err := c.RedditClient.Listing(context.Background(), subreddit, "")
	if err != nil {
		return nil, err

	}
	var postWithHighestLikes *reddit.Post
	var highestLikes int
	for _, post := range posts {
		if post.Upvotes > highestLikes {
			postWithHighestLikes = post
			highestLikes = post.Upvotes

		}

	}
	return postWithHighestLikes, nil

}
