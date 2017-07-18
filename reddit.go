// Package reddit is a simple reddit client to get resutls from a subreddit.
package reddit

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Item is a single subreddit item
type Item struct {
	Title    string
	URL      string
	Comments int `json:"num_comments"`
}

// String returns the Item represented as a string
func (i Item) String() string {
	com := ""
	switch i.Comments {
	case 0:
	case 1:
		com = " (1 comment)"
	default:
		com = fmt.Sprintf(" (%d comments)", i.Comments)
	}
	return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}

type redditClient struct {
	http.Client
}

func (c *redditClient) get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-agent", "Go Reddit/1.0")
	return c.Client.Do(req)
}

type response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

// Get returns the top results from the given subreddit. If
// there are any errors along the way, the Items will be nil
// and the error value set.
func Get(reddit string) ([]Item, error) {
	c := new(redditClient)
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	resp, err := c.get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(response)
	if err = json.NewDecoder(resp.Body).Decode(r); err != nil {
		return nil, err
	}
	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil
}
