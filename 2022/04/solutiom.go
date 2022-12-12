package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Section struct {
	start int
	end   int
}

func (s Section) contains(otherSection Section) bool {
	return s.start <= otherSection.start && s.end >= otherSection.end
}

type Pair struct {
	firstAssignment  Section
	secondAssignment Section
}

func (p Pair) sectionContained() bool {
	return p.firstAssignment.contains(p.secondAssignment) || p.secondAssignment.contains(p.firstAssignment)
}

func (p Pair) OverlappingSections() bool {
	return p.firstAssignment.start >= p.secondAssignment.start && p.firstAssignment.start <= p.secondAssignment.end ||
		p.firstAssignment.start <= p.secondAssignment.start && p.firstAssignment.end >= p.secondAssignment.start
}

const lineRegex = "([0-9]+)-([0-9]+),([0-9]+)-([0-9]+)"

func PairAssignmetnFromLine(line string) (Pair, error) {
	re := regexp.MustCompile(lineRegex)
	match := re.FindStringSubmatch(line)
	if match == nil {
		return Pair{}, errors.New("Line: " + line + " didn't match the expected fromat of " + lineRegex)
	}

	startOne, _ := strconv.Atoi(match[1])
	endOne, _ := strconv.Atoi(match[2])
	startTwo, _ := strconv.Atoi(match[3])
	endTwo, _ := strconv.Atoi(match[4])

	return Pair{Section{startOne, endOne}, Section{startTwo, endTwo}}, nil
}

func ParseFile(filePath string) ([]Pair, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []Pair{}, err
	}

	var results []Pair
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		pair, err := PairAssignmetnFromLine(line)
		if err != nil {
			return []Pair{}, err
		}

		results = append(results, pair)
	}

	return results, nil
}

func firstStar(pairs []Pair) (containedPairs int) {
	for _, pair := range pairs {
		if pair.sectionContained() {
			containedPairs++
		}
	}
	fmt.Println("First star: ", containedPairs)
	return containedPairs
}

func secondStar(pairs []Pair) (overlappingPairs int) {
	for _, pair := range pairs {
		if pair.OverlappingSections() {
			overlappingPairs++
		}
	}

	fmt.Println("Second Srat: ", overlappingPairs)
	return overlappingPairs
}

func main() {
	results, err := ParseFile("input.data")
	if err != nil {
		fmt.Println("Failed to parse input, error: ", err)
		return
	}

	firstStar(results)
	secondStar(results)
}
