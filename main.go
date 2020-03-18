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

	// Compile regex expressions
	registerRE := `^\w{2,24}`
	err = ls.ProcessLines(registerRE)
	if err != nil {
		log.Fatalf("Failed to process lines using specified regexp: %s", registerRE)
	}

	// Read User registration file
	err = ls.ReadLines("data/tweet.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	// Compile regex expressions
	tweetRE := `^\w{2,24}>\s`
	err = ls.ProcessLines(tweetRE)
	if err != nil {
		log.Fatalf("Failed to process lines using specified regexp: %s", tweetRE)
	}

	// Read User registration file
	err = ls.ReadLines("data/user.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	// Compile follow regex expressions
	followRE := `^\w{2,24}\sfollows\s`
	err = ls.ProcessLines(followRE)
	if err != nil {
		log.Fatalf("Failed to process lines using specified regexp: %s", followRE)
	}

}
