package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during file lecture: ", err)
		return
	}

	word := "XMAS"
	count := countWord(grid, word)
	fmt.Println("XMAS appears: ", count)
}

func countWord(grid [][]rune, word string) int {

	rows := len(grid)
	cols := len(grid[0])
	wordRunes := []rune(word)
	count := 0

	directions := [][2]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
	}

	for r := 0; r < rows; r++ {

		for c := 0; c < cols; c++ {

			for _, dir := range directions {

				if checkWord(grid, wordRunes, r, c, dir[0], dir[1]) {

					count++
				}
				if checkWord(grid, reverseRunes(wordRunes), r, c, dir[0], dir[1]) {

					count++
				}
			}
		}
	}
	return count
}

func checkWord(grid [][]rune, word []rune, r, c, dr, dc int) bool {

	rows := len(grid)
	cols := len(grid[0])
	wordLen := len(word)

	for i := 0; i < wordLen; i++ {

		nr := r + i*dr
		nc := c + i*dc

		if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != word[i] {
			return false
		}
	}
	return true
}

func reverseRunes(runes []rune) []rune {

	reversed := make([]rune, len(runes))
	for i, r := range runes {

		reversed[len(runes)-1-i] = r
	}
	return reversed
}
