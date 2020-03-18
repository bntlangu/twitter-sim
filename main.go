package main

import (
	"log"
	"twitter-sim/api"
	"twitter-sim/inout"
)

func main() {
	// Read User registration file
	lines, err := inout.LineStreamer("data/register.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	for _, username := range lines {

		usr, err := api.RegisterUser(username)
		if err != nil {
			log.Fatalf("Failed to open file: %s", err)
		}

		println(usr)
	}
}
