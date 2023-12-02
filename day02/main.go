package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	numberRegex := regexp.MustCompile(`\d+`)
	redRegex := regexp.MustCompile(`\d+ red`)
	greenRegex := regexp.MustCompile(`\d+ green`)
	blueRegex := regexp.MustCompile(`\d+ blue`)

	sum := 0
	sum2 := 0.

	for fileScanner.Scan() {
		line := fileScanner.Text()

		splitted := strings.Split(line, ":")
		gameId := splitted[0]
		gameData := splitted[1]

		id := numberRegex.FindString(gameId)

		rounds := strings.Split(gameData, ";")

		isValid := true

		maxBlue, maxGreen, maxRed := 0., 0., 0.

		for _, round := range rounds {
			if redRegex.MatchString(round) {
				found := redRegex.FindString(round)
				numStr := numberRegex.FindString(found)

				num, _ := strconv.Atoi(numStr)

				maxRed = math.Max(float64(maxRed), float64(num))

				if num > 12 {
					isValid = false
				}
			}

			if greenRegex.MatchString(round) {
				found := greenRegex.FindString(round)
				numStr := numberRegex.FindString(found)

				num, _ := strconv.Atoi(numStr)
				maxGreen = math.Max(float64(maxGreen), float64(num))

				if num > 13 {
					isValid = false
				}
			}

			if blueRegex.MatchString(round) {
				found := blueRegex.FindString(round)
				numStr := numberRegex.FindString(found)

				num, _ := strconv.Atoi(numStr)
				maxBlue = math.Max(float64(maxBlue), float64(num))

				if num > 14 {
					isValid = false
				}
			}
		}

		sum2 += maxRed * maxBlue * maxGreen

		if isValid {
			numId, _ := strconv.Atoi(id)
			sum += numId
		}

		// fmt.Println(maxRed, maxGreen, maxBlue)
		// fmt.Println(gameId, "dd", id, isValid, sum, sum2)
		// fmt.Scanln()
	}

	fmt.Println(sum)
	fmt.Println(sum2)
	readFile.Close()
}
