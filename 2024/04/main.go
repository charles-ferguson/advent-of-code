package main

import (
	"bufio"
	"fmt"
	"os"
)

func searchLeft(puzzle [][]rune, word string, i, j int) bool {
	if j-len(word) < -1 {
		return false
	}
	for index, letter := range word {
		if puzzle[i][j-index] != letter {
			return false
		}
	}
	return true
}

func searchRight(puzzle [][]rune, word string, i, j int) bool {
	if j+len(word) > len(puzzle[i]) {
		return false
	}
	for index, letter := range word {
		if puzzle[i][j+index] != letter {
			return false
		}
	}
	return true
}

func searchUp(puzzle [][]rune, word string, i, j int) bool {
	if i-len(word) < -1 {
		return false
	}
	for index, letter := range word {
		if puzzle[i-index][j] != letter {
			return false
		}
	}
	return true
}

func searchDown(puzzle [][]rune, word string, i, j int) bool {
	if i+len(word) > len(puzzle) {
		return false
	}
	for index, letter := range word {
		if puzzle[i+index][j] != letter {
			return false
		}
	}
	return true
}

func searchDiagonalUpLeft(puzzle [][]rune, word string, i, j int) bool {
	if i-len(word) < -1 || j-len(word) < -1 {
		return false
	}
	for index, letter := range word {
		if puzzle[i-index][j-index] != letter {
			return false
		}
	}
	return true
}

func searchDiagonalUpRight(puzzle [][]rune, word string, i, j int) bool {
	if i-len(word) < -1 || j+len(word) > len(puzzle[i]) {
		return false
	}
	for index, letter := range word {
		if puzzle[i-index][j+index] != letter {
			return false
		}
	}
	return true
}

func searchDiagonalDownLeft(puzzle [][]rune, word string, i, j int) bool {
	if i+len(word) > len(puzzle) || j-len(word) < -1 {
		return false
	}
	for index, letter := range word {
		if puzzle[i+index][j-index] != letter {
			return false
		}
	}
	return true
}

func searchDiagonalDownRight(puzzle [][]rune, word string, i, j int) bool {
	if i+len(word) > len(puzzle) || j+len(word) > len(puzzle[i]) {
		return false
	}
	for index, letter := range word {
		if puzzle[i+index][j+index] != letter {
			return false
		}
	}
	return true
}

func searchWord(puzzle [][]rune, word string) int {
	count := 0
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == rune(word[0]) {
				if searchLeft(puzzle, word, i, j) {
					count++
				}
				if searchRight(puzzle, word, i, j) {
					count++
				}
				if searchUp(puzzle, word, i, j) {
					count++
				}
				if searchDown(puzzle, word, i, j) {
					count++
				}
				if searchDiagonalUpLeft(puzzle, word, i, j) {
					count++
				}
				if searchDiagonalUpRight(puzzle, word, i, j) {
					count++
				}
				if searchDiagonalDownLeft(puzzle, word, i, j) {
					count++
				}
				if searchDiagonalDownRight(puzzle, word, i, j) {
					count++
				}
			}
		}
	}
	return count
}

func parsePuzzle(filepath string) [][]rune {
	file, err := os.Open(filepath)
	if err != nil {
		return nil
	}

	puzzle := make([][]rune, 0)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		runeLine := []rune(line)
		puzzle = append(puzzle, runeLine)
		fmt.Println(runeLine)
	}
	return puzzle
}

func printPuzzle(puzzle [][]rune) {
	for _, line := range puzzle {
		fmt.Println(string(line))
	}
}

func main() {
	puzzle := parsePuzzle("input.data")
	fmt.Println(searchWord(puzzle, "XMAS"))
	printPuzzle(puzzle)
}
