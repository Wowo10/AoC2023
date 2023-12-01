package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
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
		size := len(line)
		num := ""

		for i := 0; i < size; i++ {
			r, _ := utf8.DecodeRune([]byte{line[i]})
			if unicode.IsDigit(r) {
				num += string(r)
				break
			}
		}

		for i := size - 1; i > 0; i-- {
			r, _ := utf8.DecodeRune([]byte{line[i]})
			if unicode.IsDigit(r) {
				num += string(r)
				break
			}
		}

		converted, _ := strconv.Atoi(num)

		sum += converted

		// fmt.Println(line, num, sum)

		// fmt.Scanln()
	}

	fmt.Println(sum)

	readFile.Close()
}
