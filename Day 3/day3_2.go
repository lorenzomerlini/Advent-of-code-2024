package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()
	content, _ := os.ReadFile("input.txt")
	input := string(content)

	fmt.Println("Part 2: ", part2(input))
}

func part2(s string) int {

	return Mul(removeDont(s))

}

func removeDont(s string) string {

	var filteredDos []string
	dos := strings.Split(s, "do()")
	for _, do := range dos {
		part, _, _ := strings.Cut(do, "don't()")
		filteredDos = append(filteredDos, part)
	}
	return strings.Join(filteredDos, "")
}

func Mul(s string) int {

	re := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	matches := re.FindAllStringSubmatch(s, -1)
	sum := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += (x * y)
	}
	return sum
}
