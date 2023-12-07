package main

import (
	"strconv"
	"strings"
)

type Category int64

const (
	Five     Category = 20
	Four     Category = 12
	Full     Category = 8
	Three    Category = 6
	TwoPairs Category = 4
	Pair     Category = 2
	None     Category = 0
)

func HigherCategory(i int) int {
	if i == 0 {
		return 2
	}

	if i == 2 {
		return 6
	}

	if i == 4 {
		return 8
	}

	if i == 6 {
		return 12
	}

	if i == 8 {
		return 12
	}

	if i == 12 {
		return 20
	}

	return i
}

var Order []rune = []rune{
	'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J',
}

func FindCardValue(r rune) int {
	for i := 0; i < len(Order); i++ {
		if r == Order[i] {
			return i
		}
	}
	return -1
}

type Hand struct {
	Cards    string
	Category Category
	Bid      int
}

func CmpHands(a, b Hand) int {
	if a.Category > b.Category {
		return 1
	}

	if a.Category < b.Category {
		return -1
	}

	for i := 0; i < len(a.Cards); i++ {
		firstCard := a.Cards[i]
		secondCard := b.Cards[i]

		if firstCard != secondCard {
			if FindCardValue(rune(firstCard)) < FindCardValue(rune(secondCard)) {
				return 1
			} else {
				return -1
			}
		}
	}

	return 0
}

func CheckCategory(s string) Category {
	count := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && i != j && s[i] != 'J' {
				count++
			}
		}
	}

	jokers := CountJokers(s)

	for i := 0; i < jokers; i++ {
		count = HigherCategory(count)
	}

	return Category(count)
}

func CountJokers(s string) (count int) {
	for _, r := range s {
		if r == 'J' {
			count++
		}
	}

	return
}

func CreateHand(s string) Hand {
	split := strings.Split(s, " ")

	i, _ := strconv.Atoi(split[1])

	return Hand{
		Cards:    split[0],
		Bid:      i,
		Category: CheckCategory(split[0]),
	}
}
