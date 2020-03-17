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

// Lookup tweet test

// Add a comment to a tweet test
