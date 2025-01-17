package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TRAIL_START = 0
const TRAIL_END = 9

func pointInBounds(p *Point, grid [][]int) bool {
	return p.i >= 0 && p.i < len(grid) && p.j >= 0 && p.j < len(grid[0])
}

type Point struct{ i, j int }

var DELTA = [...]Point{{i: -1, j: 0}, {i: 1, j: 0}, {i: 0, j: -1}, {i: 0, j: 1}}

// Basically a simple dfs
func findDistinctTrails(grid [][]int, start *Point) int {
	startHeight := grid[start.i][start.j]
	if startHeight == TRAIL_END {
		return 1
	}

	trails := 0
	for _, delta := range DELTA {
		neighbor := Point{i: start.i + delta.i, j: start.j + delta.j}
		if !pointInBounds(&neighbor, grid) {
			continue
		}

		neighborHeight := grid[neighbor.i][neighbor.j]
		if neighborHeight-startHeight == 1 {
			neighborTrails := findDistinctTrails(grid, &neighbor)
			trails += neighborTrails
		}
	}

	return trails
}

func main() {
	bio := bufio.NewScanner(os.Stdin)

	grid := [][]int{}
	for bio.Scan() {
		line := bio.Text()
		rawDigits := strings.Split(line, "")
		row := make([]int, len(rawDigits))
		for i, rawDigit := range rawDigits {
			row[i], _ = strconv.Atoi(rawDigit)
		}
		grid = append(grid, row)
	}

	totalTrails := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == TRAIL_START {
				trails := findDistinctTrails(grid, &Point{i, j})
				totalTrails += trails
			}
		}
	}

	fmt.Println(totalTrails)
}
