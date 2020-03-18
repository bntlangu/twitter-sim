package inout

import (
	"bufio"
	"errors"
	"os"
)

// LineStreamer reads an input file line by line and returns
// the corresponding list of strings.
func LineStreamer(filename string) ([]string, error) {
	readFile, err := os.Open(filename)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()
	return fileTextLines, nil
}
