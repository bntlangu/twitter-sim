package model

// User is the data struct for housing a user
type User struct {
	*Model
	User    string           // User's name
	Tweets  map[string]Tweet // Tweets by this user
	Follows []string         // Users that this user follows
}
