package main

import (
	"fmt"
	"os"
	"strings"
)

var board []string

type direction int

const (
	up    direction = 0
	down  direction = 1
	left  direction = 2
	right direction = 3
)

func main() {
	readFile, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	board = strings.Split(string(readFile), "\n")

	startX, startY := 0, 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == 'S' {
				startX = i
				startY = j

				fmt.Println("start:", startX, startY)
			}
		}
	}

	path := followPath(startX-1, startY, up)

	if path != 0 {
		fmt.Println(path / 2)
		return
	}

	path = followPath(startX, startY-1, left)
	if path != 0 {
		fmt.Println(path / 2)
		return
	}
	path = followPath(startX, startY+1, right)
	if path != 0 {
		fmt.Println(path / 2)
		return
	}
	path = followPath(startX+1, startY, down)
	if path != 0 {
		fmt.Println(path / 2)
		return
	}

	panic("No answer")
}

func followPath(x, y int, direction direction) (path int) {
	for {
		if x >= len(board) || y >= len(board[x]) || x < 0 || y < 0 {
			return 0
		}

		sign := board[x][y]

		switch sign {
		case 'S':
			return path + 1
		case '.':
			return 0
		case '|':
			if direction == left || direction == right {
				return 0
			}
			if direction == up {
				x--
			}
			if direction == down {
				x++
			}
		case '-':
			if direction == up || direction == down {
				return 0
			}
			if direction == left {
				y--
			}
			if direction == right {
				y++
			}
		case '7':
			if direction == left || direction == down {
				return 0
			}
			if direction == up {
				direction = left
				y--
			}
			if direction == right {
				direction = down
				x++
			}
		case 'J':
			if direction == left || direction == up {
				return 0
			}
			if direction == right {
				direction = up
				x--
			}
			if direction == down {
				direction = left
				y--
			}
		case 'F':
			if direction == right || direction == down {
				return 0
			}
			if direction == up {
				direction = right
				y++
			}
			if direction == left {
				direction = down
				x++
			}
		case 'L':
			if direction == right || direction == up {
				return 0
			}
			if direction == down {
				direction = right
				y++
			}
			if direction == left {
				direction = up
				x--
			}
		}

		path++
	}
}
