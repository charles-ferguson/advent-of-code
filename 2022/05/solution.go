package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CrateStacks struct {
	stacks [][]rune
}

const stackWidth = 4
const stackCrateIndex = 1

func (cs CrateStacks) Print() {
	maxStackHeight := 0
	for _, stack := range cs.stacks {
		if len(stack) > maxStackHeight {
			maxStackHeight = len(stack)
		}
	}
	fmt.Println("Current Stack")

	for index := maxStackHeight - 1; index >= 0; index-- {
		for _, stack := range cs.stacks {
			if len(stack)-1 < index {
				fmt.Print("    ")
			} else {
				fmt.Print("[", string(stack[index]), "] ")
			}
		}
		fmt.Print("\n")
	}

}
func ParseCrateStacks(chunk []string) CrateStacks {
	numberedStacksStrings := chunk[len(chunk)-1]
	splitStackNumbers := strings.Fields(numberedStacksStrings)
	numberOfStacks := len(splitStackNumbers)

	var stacks [][]rune
	for index := 0; index < numberOfStacks; index++ {
		stacks = append(stacks, []rune{})
	}

	for chunkIndex := len(chunk) - 2; chunkIndex >= 0; chunkIndex-- {
		for stackIndex := 0; stackIndex < numberOfStacks; stackIndex++ {
			crate := []rune(chunk[chunkIndex])[stackIndex*stackWidth+stackCrateIndex]
			if crate != ' ' {
				stacks[stackIndex] = append(stacks[stackIndex], crate)
			}
		}
	}

	return CrateStacks{stacks}
}
func (cs CrateStacks) TopOfStacks() (tops []rune) {
	for _, stack := range cs.stacks {
		tops = append(tops, stack[len(stack)-1])
	}
	return
}

type Move struct {
	fromStack      int
	numberOfCrates int
	toStack        int
}

func ParseMoveFromLine(line string) Move {
	re := regexp.MustCompile(
		"move (?P<numberOfCrates>[0-9]+) from (?P<fromStack>[0-9]+) to (?P<toStack>[0-9]+)")
	match := re.FindStringSubmatch(line)
	if match != nil {
		numberOfStacks, _ := strconv.Atoi(match[1])
		fromStack, _ := strconv.Atoi(match[2])
		toStack, _ := strconv.Atoi(match[3])
		return Move{fromStack - 1, numberOfStacks, toStack - 1}
	}
	return Move{}
}

func ParseFile(filepath string) (CrateStacks, []Move, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return CrateStacks{}, []Move{}, err
	}

	var lines []string
	var cratesStacks CrateStacks
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		if len(fileScanner.Text()) <= 1 {
			cratesStacks = ParseCrateStacks(lines)
			break
		}
		lines = append(lines, fileScanner.Text())
	}

	moves := []Move{}
	for fileScanner.Scan() {
		moves = append(moves, ParseMoveFromLine(fileScanner.Text()))
	}

	return cratesStacks, moves, nil
}

func (cs CrateStacks) MoveStack(move Move) error {
	if move.fromStack+1 > len(cs.stacks) {
		return errors.New(fmt.Sprintf("Can't move crates from stack: %d there only %d stacks.", move.fromStack, len(cs.stacks)))
	}
	if move.toStack+1 > len(cs.stacks) {
		return errors.New(fmt.Sprintf("Can't move crates to stack: %d there only %d stacks.", move.fromStack, len(cs.stacks)))
	}
	if move.numberOfCrates > len(cs.stacks[move.fromStack]) {
		return errors.New(fmt.Sprintf(
			"Can't move %d crates from stack: %d there only %d crates.",
			move.numberOfCrates,
			move.fromStack,
			len(cs.stacks[move.fromStack]),
		))
	}

	moveTopIndex := len(cs.stacks[move.fromStack])
	moveBottomIndex := moveTopIndex - move.numberOfCrates
	crates, fromStack := cs.stacks[move.fromStack][moveBottomIndex:moveTopIndex], cs.stacks[move.fromStack][0:moveBottomIndex]
	cs.stacks[move.toStack] = append(cs.stacks[move.toStack], crates...)
	cs.stacks[move.fromStack] = fromStack

	return nil
}
func (cs CrateStacks) Move(move Move) error {
	if move.fromStack+1 > len(cs.stacks) {
		return errors.New(fmt.Sprintf("Can't move crates from stack: %d there only %d stacks.", move.fromStack, len(cs.stacks)))
	}
	if move.toStack+1 > len(cs.stacks) {
		return errors.New(fmt.Sprintf("Can't move crates to stack: %d there only %d stacks.", move.fromStack, len(cs.stacks)))
	}
	if move.numberOfCrates > len(cs.stacks[move.fromStack]) {
		return errors.New(fmt.Sprintf(
			"Can't move %d crates from stack: %d there only %d crates.",
			move.numberOfCrates,
			move.fromStack,
			len(cs.stacks[move.fromStack]),
		))
	}

	for index := 0; index < move.numberOfCrates; index++ {
		crate, fromStack := cs.stacks[move.fromStack][len(cs.stacks[move.fromStack])-1], cs.stacks[move.fromStack][:len(cs.stacks[move.fromStack])-1]
		cs.stacks[move.toStack] = append(cs.stacks[move.toStack], crate)
		cs.stacks[move.fromStack] = fromStack
	}
	return nil
}

func firstStar(crateStacks CrateStacks, moves []Move) string {
	for _, move := range moves {
		err := crateStacks.Move(move)
		if err != nil {
			fmt.Println(err)
		}
		// crateStacks.Print()
	}

	tops := string(crateStacks.TopOfStacks())
	fmt.Println("First Star: ", tops)
	return tops
}

func secondStar(crateStacks CrateStacks, moves []Move) string {
	for _, move := range moves {
		err := crateStacks.MoveStack(move)
		if err != nil {
			fmt.Println(err)
		}
		// crateStacks.Print()
	}
	tops := string(crateStacks.TopOfStacks())

	fmt.Println("Second Star: ", tops)
	return tops
}

func main() {
	crateStacks, moves, _ := ParseFile("input.data")
	firstStar(crateStacks, moves)
	crateStacks, moves, _ = ParseFile("input.data")
	secondStar(crateStacks, moves)
}
