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

// ProcessLines parses the lines stored in the LineStreamer
// and attempts to process them according to their content
// using regex expressions to match each line to an appropriate function call.
func (ls *LineStreamer) ProcessLines(re string) error {
	for _, line := range ls.Lines {
		regexRE := regexp.MustCompile(re)

		// Process line to determine what action must be taken
		regexMatch := regexRE.FindString(line)

		if regexMatch != "" {
			fmt.Println(regexMatch)
			if _, err := api.RegisterUser(regexMatch); err != nil {
				return errors.New(err.Error())
			}
		}
	}

	return nil
}
