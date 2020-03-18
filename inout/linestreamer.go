package inout

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"twitter-sim/api"
)

// LineStreamer is a struct for performing file input/out and processing functions
type LineStreamer struct {
	Filename string
	Lines    []string
}

// REGIST, TWEET, and FOLLOW are constants enumerating input stream types
const (
	REGISTER = iota
	TWEET
	FOLLOW
)

// ReadLines reads an input file line by line and returns
// the corresponding list of strings.
func (ls *LineStreamer) ReadLines(filename string) error {
	ls.Filename = filename
	readFile, err := os.Open(filename)
	if err != nil {
		return errors.New(err.Error())
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		ls.Lines = append(ls.Lines, fileScanner.Text())
	}

	readFile.Close()
	return nil
}

func processRegister(line string) (string, error) {

	// Compile regex expressions
	registerRE := regexp.MustCompile(`^\w{2,24}`)
	// Process line to determine what action must be taken
	regexMatch := registerRE.FindString(line)

	if regexMatch != "" {
		fmt.Println(regexMatch)
		username, err := api.RegisterUser(regexMatch)
		if err != nil {
			return "", errors.New(err.Error())
		}

		return username, nil
	}

	return "", errors.New("Failed to process user registeration for " + line)
}

func processTweet(line string) (string, error) {

	// Compile regex expressions
	tweetRE := regexp.MustCompile(`>\s`)
	// Process line to determine what action must be taken
	regexMatch := tweetRE.Split(line, -1)

	if len(regexMatch) > 1 {
		fmt.Println(regexMatch)
		username, err := api.Tweet(regexMatch[0], regexMatch[1])
		if err != nil {
			return "", errors.New(err.Error())
		}

		return username, nil
	}

	return "", errors.New("Failed to process user registeration for " + line)
}

func processFollow(line string) (string, error) {

	// Compile regex expressions
	followRE := regexp.MustCompile(`\sfollows\s`)
	// Split according to the expression specified
	regexMatch := followRE.Split(line, -1)

	if len(regexMatch) > 1 {
		fmt.Println(regexMatch)
		username, err := api.Follow(regexMatch[0], regexMatch[1])
		if err != nil {
			return "", errors.New(err.Error())
		}

		return username, nil
	}

	return "", errors.New("Failed to process user registeration for " + line)
}

// ProcessLines parses the lines stored in the LineStreamer
// and attempts to process them according to their content
// using regex expressions to match each line to an appropriate function call.
func (ls *LineStreamer) ProcessLines(inputType int) (err error) {
	for _, line := range ls.Lines {
		switch inputType {
		case REGISTER:
			_, err = processRegister(line)
			break
		case TWEET:
			_, err = processTweet(line)
			break
		case FOLLOW:
			_, err = processFollow(line)
			break
		}
	}

	return nil
}
