package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// Open input.txt
	file, _ := os.Open("input.txt")
	defer file.Close()

	var col1, col2 []int // two slices for every column

	// read
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) < 2 {
			continue
		}
		// converts number
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])

		col1 = append(col1, num1)
		col2 = append(col2, num2)
	}

	// sort numbers
	sort.Ints(col1)
	sort.Ints(col2)

	var result int

	for i := 0; i < len(col1); i++ {

		if col1[i] > col2[i] {

			r1 := col1[i] - col2[i]
			result += r1
		}
		if col1[i] < col2[i] {

			r1 := col2[i] - col1[i]
			result += r1
		}
		if col1[i] == col2[i] {

			result += 0
		}
	}

	fmt.Println("Result: ", result)

	Part2(col1, col2)
}

func Part2(col1 []int, col2 []int) {

	var similarity int
	for _, num1 := range col1 {

		for _, num2 := range col2 {

			if num1 == num2 {

				similarity += num1
			}
		}
	}
	fmt.Println("Similarity: ", similarity)
}
