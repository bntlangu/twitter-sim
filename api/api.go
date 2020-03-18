package api

import (
	"errors"
	"twitter-sim/model"
)

// DB is the object which will persist data for the lifetime of the application.
var DB RAMdb

func init() {
	DB.Initialise()
}

// RegisterUser is the API call for adding users to the application
func RegisterUser(username string) (string, error) {
	// <TODO>
	// 1. Check if user exists
	// 2. Construct new user
	// 2. Save new user to DB
	// 3. Handle errors

	return "", errors.New("api error: user could not be registered")
}

// Tweet adds a tweet to the list of tweets belonging to a user
func Tweet(input []byte) (string, error) {

	// <TODO>
	// 1. Append to user's tweet list
	// 2. Save user back to DB
	// 3. Handle errors
	return "", errors.New("api error: tweet failed")
}

// Comment add a comment to the list of comments attached to a tweet
func Comment(username string, tweetID string, input string) (*model.Tweet, error) {

	// <TODO>

	// 1. Append to user's tweet list
	// 2. Save user back to DB
	// 3. Handle errors
	return nil, errors.New("api error: comment could not be added to tweet")
}

// Follow adds another user to the list of users followed by the calling user.
func Follow(inputUser string, followedUser string) error {

	// <TODO>
	// 1. Append to user's following list
	// 2. Save user back to DB
	// 3. Handle errors
	return nil
}
