package main

import (
	"bufio"
	"fmt"
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
	vertLoc, widthLoc := 1, 3
	// Assumes same width on all lines
	maxWidth := len(fileTextLines[0])

	treeCount := 0

	for i, rowData := range fileTextLines {
		if i == vertLoc {
			locItem := rowData[widthLoc-1]
			fmt.Printf("%s - %d, %d (%c)\n", rowData, vertLoc, widthLoc, locItem)
			if locItem == '#' {
				treeCount++
			}
			vertLoc = vertLoc + 1
			widthLoc = widthLoc + 3
			if widthLoc > maxWidth {
				widthLoc = widthLoc - maxWidth
			}
		}
	}
	fmt.Printf("Total Trees: %d\n", treeCount)
}
