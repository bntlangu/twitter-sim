package api

import (
	"errors"
	"twitter-sim/db"
	"twitter-sim/model"
)

// DB is the object which will persist data for the lifetime of the application.
var DB *db.RAMdb

// RegisterUser is the API call for adding users to the application
func RegisterUser(username string) (*model.User, error) {
	// <TODO>
	// 1. Check if user exists
	// 2. Construct new user
	// 2. Save new user to DB
	// 3. Handle errors

	return nil, errors.New("api error: user could not be registered")
}

// LookupUser is a utility that performs a lookup of users by their username
func LookupUser(username string) (*model.User, error) {

	if user := DB.LookupUser(username); user != nil {
		return user, nil
	}
	// <TODO>
	// 1. Search DB for user with username equal to input string
	// 2. Return user if found, nil if not.
	// 3. Handle errors

	return nil, errors.New("api error: user not found")
}

// Tweet adds a tweet to the list of tweets belonging to a user
func Tweet(input []byte) (*model.Tweet, error) {

	// <TODO>
	// 1. Append to user's tweet list
	// 2. Save user back to DB
	// 3. Handle errors
	return nil, errors.New("api error: tweet failed")
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
