package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	sum := 0

	hands := []Hand{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		hands = append(hands, CreateHand(line))
	}

	slices.SortFunc(hands, CmpHands)

	for i, hand := range hands {
		sum += hand.Bid * (i + 1)
	}

	fmt.Println(sum)
}
