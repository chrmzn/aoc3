package main

import (
	"bufio"
	"fmt"
	"os"
)

func runRouteAnalysis(route []string, vertMov int, widthMov int) (treeCount int) {
	// Starting at 0,0
	vertLoc, widthLoc := 0+vertMov, 1+widthMov
	// Assumes same width on all lines
	maxWidth := len(route[0])

	for i, rowData := range route {
		if i == vertLoc {
			locItem := rowData[widthLoc-1]
			// fmt.Printf("%s - %d, %d (%c)\n", rowData, vertLoc, widthLoc, locItem)
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

	return
}

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

	runTypes := [][]int{{1, 1}, {1, 3}}

	for _, run := range runTypes {
		vertMov, widthMov := run[0], run[1]
		treeCount := runRouteAnalysis(fileTextLines, vertMov, widthMov)
		fmt.Printf("Total Trees (%d, %d): %d\n", vertMov, widthMov, treeCount)
	}
}
