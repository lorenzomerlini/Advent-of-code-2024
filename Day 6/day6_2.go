package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y      int
	direction int
}

type Coordinates struct {
	x, y int
}

const (
	up    = 0
	down  = 1
	right = 2
	left  = 3
)

func main() {

	grid := loadMap("input.txt")
	start := findStartingPoint(grid)

	pathLength, path := findPath(grid, start)
	loops := CreateLoop(grid, start, path)

	fmt.Println("Lunghezza percorso: ", pathLength)
	fmt.Println("Numero loops: ", loops)
}

func findStartingPoint(grid [][]rune) Point {

	directionMap := map[rune]int{'^': up, 'v': down, '>': right, '<': left}
	for y, row := range grid {
		for x, cell := range row {
			if dir, exists := directionMap[cell]; exists {

				return Point{x, y, dir}
			}
		}
	}
	panic("Punto di partenza non trovato")
}

func findPath(grid [][]rune, start Point) (int, map[Coordinates]int) {

	path := make(map[Coordinates]int)
	steps := 0
	current := start

	for {
		coord := Coordinates{current.x, current.y}
		if _, visited := path[coord]; !visited {
			steps++
			path[coord] = current.direction
		}

		valid, next := nextStep(grid, current)
		if !valid {
			break
		}
		current = next
	}
	return steps, path
}

func CreateLoop(grid [][]rune, start Point, path map[Coordinates]int) int {

	loopCount := 0
	for coord := range path {

		if coord.x == start.x && coord.y == start.y {
			continue
		}
		if grid[coord.y][coord.x] == '.' {
			grid[coord.y][coord.x] = '#'
			if checkLoop(grid, start) {
				loopCount++
			}
			grid[coord.y][coord.x] = '.'
		}
	}
	return loopCount
}

func checkLoop(grid [][]rune, start Point) bool {

	visited := make(map[Point]struct{})
	current := start

	for {
		if _, seen := visited[current]; seen {
			return true
		}
		visited[current] = struct{}{}
		valid, next := nextStep(grid, current)
		if !valid {
			break
		}
		current = next
	}
	return false
}

func nextStep(grid [][]rune, current Point) (bool, Point) {
	dx, dy := 0, 0
	switch current.direction {
	case up:
		dy = -1
	case down:
		dy = 1
	case right:
		dx = 1
	case left:
		dx = -1
	}
	next := Point{current.x + dx, current.y + dy, current.direction}
	if next.x < 0 || next.y < 0 || next.x >= len(grid[0]) || next.y >= len(grid) {
		return false, current
	}

	if grid[next.y][next.x] == '#' {
		return nextStep(grid, turn90(current))
	}
	return true, next
}
func turn90(current Point) Point {
	switch current.direction {
	case up:
		current.direction = right
	case right:
		current.direction = down
	case down:
		current.direction = left
	case left:
		current.direction = up
	}
	return current
}

// Carica la mappa da un file
func loadMap(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid
}
