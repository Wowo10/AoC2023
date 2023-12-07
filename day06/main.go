package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputContent, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(inputContent), "\n")

	timeLine := lines[0]
	distanceLine := lines[1]

	fmt.Println(solve(timeLine, distanceLine))

	timeLine = strings.ReplaceAll(lines[0], " ", "")
	distanceLine = strings.ReplaceAll(lines[1], " ", "")

	fmt.Println(solve(timeLine, distanceLine))
}

func solve(timeLine string, distanceLine string) int {
	timeLines := findNumbers(timeLine)
	distanceLines := findNumbers(distanceLine)

	if len(timeLines) != len(distanceLines) {
		panic("Parse f$*#ed.")
	}

	amount := 1

	for i := 0; i < len(timeLines); i++ {
		time := timeLines[i]
		distance := distanceLines[i]

		count := 0

		for i := 0; i <= time; i++ {

			if i*(time-i) > distance {
				count++
			}
		}

		amount *= count
	}

	return amount
}

func findNumbers(s string) []int {
	numberRegex := regexp.MustCompile(`\d+`)

	matches := numberRegex.FindAllString(s, -1)

	return mapSlice(matches, func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})
}

func mapSlice[T any, M any](a []T, f func(T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}
