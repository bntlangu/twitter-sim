package api

import "testing"

func TestRegisterUser(t *testing.T) {
	if _, err := RegisterUser("Tom"); err != nil {
		t.Error("Failed to register Tom")
	}

	if _, err := RegisterUser("Dick"); err != nil {
		t.Error("Failed to register Dick")
	}

	if _, err := RegisterUser("Harry"); err != nil {
		t.Error("Failed to register Harry")
	}

	if DB.GetNumUsers() != 3 {
		t.Error("Incorrect number of users found.")
	}
}

func TestTweet(t *testing.T) {
	if _, err := Tweet("Tom", "Tom's first tweet"); err != nil {
		t.Error("Failed to register Tom")
	}

	if _, err := Tweet("Dick", "Dick's first tweet"); err != nil {
		t.Error("Failed to register Dick")
	}

	if _, err := Tweet("Tom", "Tom's second tweet"); err != nil {
		t.Error("Failed to register Harry")
	}

	usr, err := DB.LookupUser("Tom")
	if err != nil {
		t.Error(err.Error())
	}

	if len(usr.Tweets) != 2 {
		t.Error("Incorrect number of tweets returned for Tom. Expected 2")
	}
}

func TestFollow(t *testing.T) {
	if _, err := Follow("Harry", "Tom"); err != nil {
		t.Error("Harry failed to follow Tom")
	}

	if _, err := Follow("Dick", "Harry"); err != nil {
		t.Error("Dick failed to follow Harry")
	}

	if _, err := Follow("Harry", "Dick"); err != nil {
		t.Error("Harry failed to follow Dick")
	}

	usr, err := DB.LookupUser("Harry")
	if err != nil {
		t.Error(err.Error())
	}
	if len(usr.Following) != 3 {
		t.Errorf("Incorrect number of users followed returned for %s. Expected %d, got %d",
			usr.User, 3, len(usr.Following))
	}
}

func TestGetFeed(t *testing.T) {

}
