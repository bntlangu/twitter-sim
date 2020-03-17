package db

import (
	"errors"
	"twitter-sim/model"
)

// RAMdb is the database object, which is comprised of a map of users.
type RAMdb struct {
	users model.UserList
}

// Initialise prepares the map to use for storing users.
func (db *RAMdb) Initialise() {
	db.users = make(model.UserList)
}

// AddUser inserts a new user into the DB if it doesn't already exist.
// Returns err otherwise.
func (db *RAMdb) AddUser(username string) error {
	// Check if user exists
	if user := db.LookupUser(username); user != nil {
		return errors.New("db error: user already exists")
	}
	// Create a user object
	var newUser model.User

	newUser.Model.SetTimestamp()
	newUser.User = username
	newUser.Tweets = make(model.TweetList)
	// Save the new user
	db.users[username] = newUser
	return nil

}

// LookupUser returns the user object associate with the specified username, nil if not found
func (db *RAMdb) LookupUser(username string) *model.User {
	user, present := db.users[username]

	if !present {
		return nil
	}

	return &user

}

// Update a user

// Add a tweet

// Lookup a tweet

// Add a comment to a tweet
