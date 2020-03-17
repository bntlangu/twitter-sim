package db

import "testing"

var DB RAMdb

func init() {
	DB.Initialise()
}

// Add user test
func TestAddUser(t *testing.T) {
	// Test positive case
	if err := DB.AddUser("TestUser"); err != nil {
		t.Error(err.Error())
	}

	// Test negative case
	if err := DB.AddUser("TestUser"); err == nil {
		t.Error("AddUser should fail, if the user already exists.")
	}

}

// Lookup user test
func TestLookupUser(t *testing.T) {
	// Test positive case
	if usr := DB.LookupUser("TestUser"); usr == nil {
		t.Error("TestUser not found in the DB though expected.")
	}

	// Test negative case
	if err := DB.LookupUser("WrongUser"); err != nil {
		t.Error("LookupUser should return nil, if the user doesn't exist.")
	}

}

// Update user test

// Add tweet test
func TestAddTweet(t *testing.T) {
	// Test positive case
	username := "TestUser"
	tweet := "This is a test tweet."
	if id, err := DB.AddTweet(username, tweet); err != nil {
		t.Error(err.Error())
	}

	// Test negative case
	username = "WrongUser"
	tweet = "This is a test tweet by the wrong user."
	if _, err := DB.AddTweet(username, tweet); err == nil {
		t.Error("AddTweet should fail, if the user doesn't exist.")
	}

}

// Lookup tweet test
func TestLookupTweet(t *testing.T) {
	// Test positive case
	username := "TestUser"
	tweet := "This is another test tweet."
	tweetID, err := DB.AddTweet(username, tweet)

	if err != nil {
		t.Error(err.Error())
	}

	// Test positive case
	if tweet := DB.LookupTweet(username, tweetID); usr == tweet {
		t.Error("Test tweet not found in the DB though expected.")
	}

}

// Add a comment to a tweet test
