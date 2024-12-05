package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rules, pages []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		rules = append(rules, scanner.Text())
	}

	for scanner.Scan() {
		pages = append(pages, scanner.Text())
	}

	var incorrectPages [][]string

	for _, page := range pages {

		list := strings.Split(page, ",")
		if Rules(list, rules) {
			continue
		} else {
			incorrectPages = append(incorrectPages, list)
		}
	}

	var sum int
	for _, page := range incorrectPages {

		order := Reorder(page, rules)
		middle, _ := strconv.Atoi(order[len(order)/2])
		sum += middle
	}
	fmt.Println("Sum: ", sum)
}

func Reorder(page []string, rules []string) []string {

	for !Rules(page, rules) {

		for _, rule := range rules {

			rule := strings.Split(rule, "|")
			indexX := Find(page, rule[0])
			indexY := Find(page, rule[1])
			if indexX == -1 || indexY == -1 {
				continue
			}
			if indexX > indexY {

				page = slices.Replace(page, indexX, indexX+1, rule[1])
				page = slices.Replace(page, indexY, indexY+1, rule[0])
			}
		}
	}
	return page
}

func Find(list []string, element string) int {
	for i, item := range list {
		if item == element {
			return i
		}
	}
	return -1
}

func Rules(page []string, rules []string) bool {

	for _, rule := range rules {

		rule := strings.Split(rule, "|")
		indexX := Find(page, rule[0])
		indexY := Find(page, rule[1])
		if indexX == -1 || indexY == -1 {
			continue
		}
		if indexX > indexY {
			return false
		}
	}
	return true
}
