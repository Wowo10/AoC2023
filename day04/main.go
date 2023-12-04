package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		splitted := strings.Split(strings.Split(line, ":")[1], "|")

		winning := strings.ReplaceAll(strings.Trim(splitted[0], " "), "  ", " ")
		mine := strings.ReplaceAll(strings.Trim(splitted[1], " "), "  ", " ")

		winningNumbers := strings.Split(winning, " ")
		mineNumbers := strings.Split(mine, " ")

		temp := 1
		for _, number := range mineNumbers {
			if contains(winningNumbers, number) {
				temp *= 2
			}
		}

		if temp != 1 {
			sum += temp / 2
		}
	}

	fmt.Println(sum)
}

func contains(arr []string, a string) bool {
	for _, e := range arr {
		if strings.EqualFold(e, a) {
			return true
		}
	}

	return false
}
