package main

import (
	"fmt"
	"log"
	"twitter-sim/api"
	"twitter-sim/inout"
)

func main() {
	// Read User registration file
	var ls inout.LineStreamer
	err := ls.ReadLines("data/register.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	err = ls.ProcessLines(inout.REGISTER)
	if err != nil {
		log.Fatalf("Failed to process lines using REGISTER")
	}

	// Read tweets file
	err = ls.ReadLines("data/tweet.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	err = ls.ProcessLines(inout.TWEET)
	if err != nil {
		log.Fatalf("Failed to process lines using specified TWEET")
	}
	// Read follow file
	err = ls.ReadLines("data/user.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	err = ls.ProcessLines(inout.FOLLOW)
	if err != nil {
		log.Fatalf("Failed to process lines using specified FOLLOW")
	}

	feed, err := api.GetFeed("Alan")
	if err == nil {
		for elem := range feed {
			fmt.Println(elem)
		}
	}

	feed, err = api.GetFeed("Martin")
	if err == nil {
		for elem := range feed {
			fmt.Println(elem)
		}
	}

	feed, err = api.GetFeed("Ward")
	if err == nil {
		for elem := range feed {
			fmt.Println(elem)
		}
	}

}
