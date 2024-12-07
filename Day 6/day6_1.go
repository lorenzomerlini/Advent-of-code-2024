package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	grid := loadMap("input.txt")
	startX, startY, startDirection := guardPosition(grid)
	//fmt.Printf("Posizione guardia: (%d, %d), Direction: %d\n", startX, startY, startDirection)
	visited := trackMovements(grid, startX, startY, startDirection)

	var spostamenti int

	for _, c := range visited {

		if c == true {
			spostamenti++
		}
	}
	fmt.Println("Spostamenti guardia: ", spostamenti)
}

func trackMovements(grid [][]rune, startX, startY, startDirection int) map[string]bool {

	visited := make(map[string]bool)
	directions := [4][2]int{
		{-1, 0}, // Su
		{0, 1},  // Destra
		{1, 0},  // giu
		{0, -1}, // sinistra
	}
	x, y := startX, startY
	direction := startDirection

	visited[fmt.Sprintf("%d,%d", x, y)] = true

	for {

		nextX := x + directions[direction][0]
		nextY := y + directions[direction][1]

		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {

			break
		}
		if grid[nextX][nextY] == '#' {

			direction = (direction + 1) % 4

		} else {

			x, y = nextX, nextY
			visited[fmt.Sprintf("%d,%d", x, y)] = true
		}
	}
	return visited
}

func guardPosition(grid [][]rune) (int, int, int) {

	directionMap := map[rune]int{
		'^': 0, //up
		'>': 1, // right
		'v': 2,
		'<': 3, // left
	}

	for x, row := range grid {
		for y, cell := range row {

			if dir, found := directionMap[cell]; found {
				return x, y, dir
			}
		}
	}
	panic("Guard not found")
}

func loadMap(filename string) [][]rune {

	file, _ := os.Open(filename)
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	return grid
}
