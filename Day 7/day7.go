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

	var sum int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		if checkLine(line) {
			stringa := strings.Split(line, ":")
			numero, _ := strconv.Atoi(strings.TrimSpace(stringa[0]))
			sum += numero
		}
	}
	fmt.Println(sum)
}

func checkLine(line string) bool {

	var numeri []int
	stringa := strings.Split(line, ":")
	target, _ := strconv.Atoi(strings.TrimSpace(stringa[0]))
	numStr := strings.Fields(strings.TrimSpace(stringa[1]))
	for _, n := range numStr {

		numero, _ := strconv.Atoi(n)
		numeri = append(numeri, numero)
	}
	if Raggiunge(numeri, target) {
		return true
	} else {
		return false
	}
}

func Raggiunge(numeri []int, target int) bool {

	return Valuta(numeri, target, 1, numeri[0])
}

func Valuta(numeri []int, target int, idx, corrente int) bool {

	if idx == len(numeri) {
		return corrente == target
	}
	if Valuta(numeri, target, idx+1, corrente+numeri[idx]) {
		return true
	}
	if Valuta(numeri, target, idx+1, corrente*numeri[idx]) {
		return true
	}
	// PARTE 2
	concatenato, _ := strconv.Atoi(fmt.Sprintf("%d%d", corrente, numeri[idx]))
	if Valuta(numeri, target, idx+1, concatenato) {
		return true
	}
	return false
}
