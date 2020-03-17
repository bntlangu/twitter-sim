package model

import "github.com/google/uuid"

// Tweet is the data struct for housing a tweet by a user
type Tweet struct {
	*Model
	ID       string
	Text     string
	User     string
	Comments []Comment
}

// Comment is a data struct for housing comments made on tweets
type Comment struct {
	*Model
	Text string
	User string
}

// SetID sets the unique ID of a tweet.
func (twt *Tweet) SetID() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	twt.ID = id.String()
	return twt.ID, nil
}
