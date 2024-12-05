package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	var rules, updates []string
	var Checkupdates bool = false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			Checkupdates = true
			continue
		}
		if Checkupdates {
			updates = append(updates, line)
		} else {
			rules = append(rules, line)
		}
	}

	var correct []string

	for _, update := range updates {

		updatePages := strings.Split(update, ",")
		isValid := true

		for _, rule := range rules {

			rulePages := strings.Split(rule, "|")

			X, Y := rulePages[0], rulePages[1]
			indexX := indexOf(updatePages, X)
			indexY := indexOf(updatePages, Y)

			if indexX != -1 && indexY != -1 && indexX > indexY {

				isValid = false
				break
			}
		}
		if isValid {

			correct = append(correct, update)
		}
	}
	var sum int
	for i := 0; i < len(correct); i++ {

		numbers := strings.Split(correct[i], ",")
		middle := len(numbers) / 2
		num, _ := strconv.Atoi(numbers[middle])
		sum += num

	}
	fmt.Println("Sum: ", sum)
}

func indexOf(slice []string, index string) int {

	for i, v := range slice {

		if v == index {
			return i
		}
	}
	return -1
}
