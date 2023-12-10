package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	const MaxInt = int(^uint(0) >> 1)

	inputContent, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(inputContent), "\n\n")

	startPath := lines[0]
	verticesData := lines[1]

	vertices := []Vertex{}
	order := []VertexOrder{}

	wordRegex := regexp.MustCompile(`\w+`)

	for _, vertexData := range strings.Split(verticesData, "\n") {
		words := wordRegex.FindAllString(vertexData, 3)

		v := Vertex{
			name:  words[0],
			left:  words[1],
			right: words[2],
		}

		vertices = append(vertices, v)

		order = append(order, VertexOrder{
			name:   words[0],
			length: MaxInt,
		})
	}

	start := findStart(vertices, startPath)

	// fmt.Println(vertices)

	setLength(&order, start.name, 0)
	// fmt.Println(order)

	slices.SortFunc(order, func(a VertexOrder, b VertexOrder) int {
		if a.length > b.length {
			return 1
		}

		if a.length < b.length {
			return -1
		}

		return 0
	})

	currentIndex := 0

	for {
		// v := findVertex(vertices, current)

		current := order[currentIndex]

		fmt.Println(current)

		if current.name == "ZZZ" {
			fmt.Println(current.length)
			fmt.Println(len(startPath))
			return
		}

		v := findVertex(vertices, current.name)

		if current.length+1 < getLength(order, v.left) {
			setLength(&order, v.left, current.length+1)
		}

		if current.length+1 < getLength(order, v.right) {
			setLength(&order, v.right, current.length+1)
		}

		slices.SortFunc(order, func(a VertexOrder, b VertexOrder) int {
			if a.length > b.length {
				return 1
			}

			if a.length < b.length {
				return -1
			}

			return 0
		})

		currentIndex++

		// fmt.Println(vertices)
		// fmt.Println(order)
		// fmt.Println(currentIndex)

		// fmt.Scanln()
	}

}

func findVertex(vertices []Vertex, name string) Vertex {
	for i, v := range vertices {
		if v.name == name {
			return vertices[i]
		}
	}
	panic("O chuj2")
}

func getLength(order []VertexOrder, name string) int {
	for _, o := range order {
		if o.name == name {
			return o.length
		}
	}

	panic("O chuj")
}

func setLength(order *[]VertexOrder, name string, length int) {
	for i := 0; i < len(*order); i++ {
		if (*order)[i].name == name {
			(*order)[i].length = length
			return
		}
	}
}

func findStart(vertices []Vertex, path string) Vertex {
	v := findVertex(vertices, "AAA")

	for _, direction := range path {

		// fmt.Println("Goign:", string(direction))

		if direction == 'R' {

			// fmt.Println("To:", v.right)
			v = findVertex(vertices, v.right)
		} else {
			// fmt.Println("To:", v.left)
			v = findVertex(vertices, v.left)
		}
	}

	return v
}
