package api

import (
	"errors"
	"fmt"
	"sort"
	"twitter-sim/db"
	"twitter-sim/model"
)

// DB is the object which will persist data for the lifetime of the application.
var DB db.RAMdb

func init() {
	DB.Initialise()
}

// RegisterUser is the API call for adding users to the application
func RegisterUser(username string) (string, error) {
	if err := DB.AddUser(username); err != nil {
		return "", errors.New(err.Error())
	}

	return username, nil
}

// Tweet adds a tweet to the list of tweets belonging to a user
func Tweet(username string, tweetText string) (string, error) {

	if _, err := DB.AddTweet(username, tweetText); err != nil {
		return "", errors.New(err.Error())
	}

	return "@" + username + ": " + tweetText, nil
}

// Follow adds another user to the list of users followed by the calling user.
func Follow(followingUser string, followedUser string) (string, error) {

	if _, err := DB.FollowUser(followingUser, followedUser); err != nil {
		return "", errors.New(err.Error())
	}

	return followingUser + " is following " + followedUser, nil

}

// GetFeed returns a list tweets which consist a user's feed.
// The feed is returned in a time sequential order.
func GetFeed(username string) ([]model.Tweet, error) {
	usr, err := DB.LookupUser(username)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	feed := []model.Tweet{}
	//fmt.Println(usr.User)
	// Go through the list of users being followed by this user.
	for _, usrName := range usr.Following {
		followedUser, err := DB.LookupUser(usrName)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		// append list of tweets

		// Go through the list of tweets corresponding to the followed user.
		for _, twt := range followedUser.Tweets {
			feed = append(feed, twt)
			//fmt.Sprintf("\t@%s: %s", followedUser.User, twt)
		}
	}

	sort.Slice(feed[:], func(i, j int) bool {
		return feed[i].Model.Timestamp.Before(feed[j].Model.Timestamp)
	})

	printFeed(username, feed)

	return feed, nil
}

func printFeed(username string, feed []model.Tweet) {
	fmt.Println(username)

	// Go through the list of sorted tweets and output
	for _, twt := range feed {

		fmt.Printf("\t@%s: %s\n", twt.User, twt.Text)
	}
}
