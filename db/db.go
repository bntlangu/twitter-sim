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
	if _, err := db.LookupUser(username); err != nil {
		// Create a user object
		var newUser model.User
		newUser.Model.SetTimestamp()
		newUser.User = username
		newUser.Tweets = make(model.TweetList)

		// Save the new user
		db.users[username] = newUser
		return nil
	}

	return errors.New("db error: user already exists")
}

// LookupUser returns the user object associate with the specified username, nil if not found
func (db *RAMdb) LookupUser(username string) (*model.User, error) {
	user, present := db.users[username]

	if !present {
		return nil, errors.New("db error: user does not exist")
	}

	return &user, nil

}

// Update a user

// AddTweet inserts a new tweet into a user's tweet list.
// Returns an error if the user doesn't exist, otherwise the
// id of the new tweet is returned.
func (db *RAMdb) AddTweet(username string, text string) (string, error) {
	// Check if user exists
	user, err := db.LookupUser(username)

	if err != nil {
		return "", errors.New(err.Error())
	}

	// Create a tweet object
	var newTweet model.Tweet
	newTweet.Model.SetTimestamp()
	id, err := newTweet.SetID()

	if err != nil {
		return "", errors.New(err.Error())
	}

	// Save tweet to user's tweet list
	newTweet.Text = text
	user.Tweets[id] = newTweet

	// Update the user's data
	db.users[username] = *user

	return id, nil

}

// LookupTweet returns the user object associated with the specified username, nil
// if not found
func (db *RAMdb) LookupTweet(username string, tweetID string) (*model.Tweet, error) {
	user, present := db.users[username]

	if !present {
		return nil, errors.New("db error: user does not exist")
	}

	tweet, present := user.Tweets[tweetID]

	if !present {
		return nil, errors.New("db error: tweet does not exist")
	}
	return &tweet, nil

}

// Add a comment to a tweet
