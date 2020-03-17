package model

import "time"

// Model is the parent data struct for twitter-sim models
type Model struct {
	Timestamp time.Time
}

// GetTimestamp returns the timestamp of the calling object.
func (mdl *Model) GetTimestamp() time.Time {
	return mdl.Timestamp
}

// SetTimestamp sets the timestmp of the object to current time.
func (mdl *Model) SetTimestamp() {
	mdl.Timestamp = time.Now()
}
