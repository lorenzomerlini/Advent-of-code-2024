package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	pattern := regexp.MustCompile(`mul\(\d+,\d+\)`)

	var matches []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		found := pattern.FindAllString(line, -1)
		matches = append(matches, found...)
	}

	var sum int

	for _, match := range matches {

		sum += Mul(match)
	}
	fmt.Println("Somma: ", sum)
}

func Mul(match string) int {

	var n1, n2 int
	fmt.Sscanf(match, "mul(%d,%d)", &n1, &n2)

	return n1 * n2
}
