package main

import (
	"log"
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

	// Read User registration file
	err = ls.ReadLines("data/tweet.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	err = ls.ProcessLines(inout.TWEET)
	if err != nil {
		log.Fatalf("Failed to process lines using specified TWEET")
	}
	// Read User registration file
	err = ls.ReadLines("data/user.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	err = ls.ProcessLines(inout.FOLLOW)
	if err != nil {
		log.Fatalf("Failed to process lines using specified FOLLOW")
	}

}
