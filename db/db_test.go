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
	if _, err := DB.LookupUser("TestUser"); err != nil {
		t.Error(err.Error())
	}

	// Test negative case
	if usr, _ := DB.LookupUser("WrongUser"); usr != nil {
		t.Error("LookupUser should return nil, if the user doesn't exist.")
	}

}

// Update user test

// Add tweet test
func TestAddTweet(t *testing.T) {
	// Test positive case
	username := "TestUser"
	tweet := "This is a test tweet."
	if _, err := DB.AddTweet(username, tweet); err != nil {
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
	tweetText := "This is another test tweet."
	tweetID, err := DB.AddTweet(username, tweetText)

	if err != nil {
		t.Error(err.Error())
	}

	// Test positive case
	tweet, err := DB.LookupTweet(username, tweetID)

	if err != nil {
		t.Error("Test tweet not found in user's tweet list though expected.")
	}

	if tweet.ID != tweetID {
		t.Error("Test tweet IDs don't match.")
	}

}

// Add user test
func TestFollowUser(t *testing.T) {
	// Setup
	actingUser := "Follower"
	targetUser := "Followed"

	if err := DB.AddUser(actingUser); err != nil {
		t.Error(err.Error())
	}
	if err := DB.AddUser(targetUser); err != nil {
		t.Error(err.Error())
	}

	// Test positive case
	numFollowed, err := DB.FollowUser(actingUser, targetUser)

	if err != nil {
		t.Error(err.Error())
	}

	if numFollowed != 1 {
		t.Errorf("Expected to be following %d user(s), but found %d", 1, numFollowed)
	}

	t.Logf("%s now following %d users", actingUser, numFollowed)

	// Test negative case
	targetUser = "WrongUser"
	numFollowed, err = DB.FollowUser(actingUser, targetUser)

	if err != nil {
		t.Log(err.Error())

	} else {
		t.Errorf("Expected %s following %s to fail since %s does not exist.",
			actingUser, targetUser, targetUser)
	}

}

// Add a comment to a tweet test
