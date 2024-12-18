package main

import (
	"bufio"
	"fmt"
	"os"
)

type Guard struct {
	x         int
	y         int
	locations [][]int
	direction string
}

func (g Guard) nextPosition() (int, int) {
	switch g.direction {
	case "up":
		return g.x, g.y - 1
	case "down":
		return g.x, g.y + 1
	case "left":
		return g.x - 1, g.y
	case "right":
		return g.x + 1, g.y
	}
	return 0, 0
}

func (g Guard) nextDirection() string {
	switch g.direction {
	case "up":
		return "right"
	case "down":
		return "left"
	case "left":
		return "up"
	case "right":
		return "down"
	}
	return ""
}

type Puzzle [][]rune

func (p Puzzle) moveGuard(guard *Guard) {
	g := guard
	if !p.inGrid(g.x, g.y) {
		return
	}
	x, y := g.nextPosition()

	if p.inGrid(x, y) && p[y][x] == '#' {
		g.direction = g.nextDirection()
	} else {
		if p.inGrid(x, y) {
			p[g.y][g.x] = 'X'
		}
		g.x = x
		g.y = y
		g.locations = append(g.locations, []int{g.x, g.y})
	}
}

func (p Puzzle) inGrid(x, y int) bool {
	if y < 0 || y >= len(p) {
		return false
	}
	if x < 0 || x >= len(p[y]) {
		return false
	}
	return true
}

func (p Puzzle) findGuard() Guard {
	for y, row := range p {
		for x, cell := range row {
			if cell == '^' {
				return Guard{x, y, [][]int{{x, y}}, "up"}
			} else if cell == 'v' {
				return Guard{x, y, [][]int{{x, y}}, "down"}
			} else if cell == '<' {
				return Guard{x, y, [][]int{{x, y}}, "left"}
			} else if cell == '>' {
				return Guard{x, y, [][]int{{x, y}}, "right"}
			}
		}
	}
	return Guard{0, 0, [][]int{}, ""}
}

func (p Puzzle) print() {
	for _, row := range p {
		fmt.Println(string(row))
	}
}

func (p Puzzle) guardLocation() (int, int) {
	for y, row := range p {
		for x, cell := range row {
			if cell == '@' {
				return x, y
			}
		}
	}
	return 0, 0
}

func ParseFile(filepath string) Puzzle {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	puzzle := make(Puzzle, 0)
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		puzzle = append(puzzle, []rune(line))
	}

	return puzzle
}

func main() {
	puzzle := ParseFile("input.data")
	guard := puzzle.findGuard()

	puzzle.print()
	count := 0
	for puzzle.inGrid(guard.x, guard.y) {
		count++
		puzzle.moveGuard(&guard)
		// fmt.Println("After move ", guard.x, guard.y, guard.direction)
		// fmt.Println("Locations ", guard.locations)
		count = 0
		uniqLocations := make([][]int, 0)
		for _, loc := range guard.locations {
			found := false

			for _, uloc := range uniqLocations {
				if uloc[0] == loc[0] && uloc[1] == loc[1] {
					found = true
					break
				}
			}

			if !found {
				count++
				uniqLocations = append(uniqLocations, loc)
			}
		}
	}
	puzzle.print()
	fmt.Println("uniqLocations ", count-1)
}
