package main

import (
	"bufio"
	"os"
)

func main() {
	readFile, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	// Starting at 0,0
	vertLoc, widthLoc := 0, 0
	// Assumes same width on all lines
	maxWidth := len(fileTextLines[0])

}
