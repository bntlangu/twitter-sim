package model

// User is the data struct for housing a user
type User struct {
	Model
	User    string    // User's name
	Tweets  TweetList // Tweets by this user
	Follows []string  // Users that this user follows
}

// UserList is the data struct, a map, used to store users.
type UserList map[string]User
