package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// open input.txt
	file, _ := os.Open("input.txt")
	defer file.Close()

	var safereports int  // store safe reports
	var safereports2 int // store real safe reports

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		numbers := strings.Fields(line)
		var report []int

		for _, numString := range numbers {

			num, _ := strconv.Atoi(numString)
			report = append(report, num)
		}
		if Safe(report) {

			safereports++
		} else {

			for i := 0; i < len(report); i++ {

				if ProblemDamper(i, report) {

					safereports2++
					break
				}
			}
		}
	}
	fmt.Println("Safe reports, part 1: ", safereports)
	fmt.Println("Safe reports, part 2: ", safereports2+safereports)
}

func ProblemDamper(i int, report []int) bool {

	if i == 0 {
		if Safe(report[i+1:]) {
			return true
		}
		return false
	}
	if i == len(report)-1 {
		if Safe(report[:len(report)-1]) {
			return true
		}
		return false
	}

	// copy
	reportCopy := make([]int, len(report))
	copy(reportCopy, report)
	var report2 []int
	report2 = append(reportCopy[:i], reportCopy[i+1:]...)

	if Safe(report2) {
		return true
	}
	return false
}

func Safe(report []int) bool {

	status := 2
	for i := 0; i < len(report)-1; i++ {

		diff := report[i+1] - report[i]
		if diff == 0 || abs(diff) > 3 {

			return false
		}
		if status == 2 {
			if report[i] < report[i+1] {
				status = 0
			} else if report[i] > report[i+1] {
				status = 1
			} else {
				return false
			}
		} else {

			if (report[i] < report[i+1] && status == 1) || report[i] > report[i+1] && status == 0 {
				return false
			}
		}
	}
	return true
}

func abs(x int) int {

	if x < 0 {
		return -x
	}
	return x
}
