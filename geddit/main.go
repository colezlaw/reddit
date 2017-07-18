package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/colezlaw/reddit"
)

var subreddit string

func init() {
	flag.StringVar(&subreddit, "s", "golang", "The subreddit to seatch")
}

func main() {
	flag.Parse()
	items, err := reddit.Get(subreddit)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
