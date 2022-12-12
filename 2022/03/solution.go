package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type RuckSack struct {
	firstContainer  string
	secondContainer string
}

func (rs RuckSack) MisplacedItem() (misplaced rune, err error) {
	inFirstContainer := make(map[rune]struct{})
	inSecondContainer := make(map[rune]struct{})
	for index := 0; index < len(rs.firstContainer); index++ {
		firstChar := rune(rs.firstContainer[index])
		inFirstContainer[firstChar] = struct{}{}
		secondChar := rune(rs.secondContainer[index])
		inSecondContainer[secondChar] = struct{}{}

		if _, ok := inFirstContainer[secondChar]; ok {
			return secondChar, nil
		}
		if _, ok := inSecondContainer[firstChar]; ok {
			return firstChar, nil
		}
	}
	return ' ', errors.New(fmt.Sprintf("There was no misplaced Item in the Rucksack %s", rs))
}

func FindBadge(ruckSacks []RuckSack) rune {
	var inRuckSacks []map[rune]struct{}
	for index, ruckSack := range ruckSacks {
		inRuckSacks = append(inRuckSacks, make(map[rune]struct{}))
		for _, char := range []rune(ruckSack.firstContainer + ruckSack.secondContainer) {
			inRuckSacks[index][char] = struct{}{}
		}
	}

	for char, _ := range inRuckSacks[0] {
		inAll := true
		for _, inRuckSack := range inRuckSacks {
			if _, included := inRuckSack[char]; !included {
				inAll = false
				break
			}
		}

		if inAll {
			return char
		}
	}
	return ' '
}

func ItemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - int('a') + 1
	}
	if item >= 'A' && item <= 'Z' {
		return int(item) - int('A') + 1 + 26
	}
	return 0
}

func ruckSackFromLine(line string) (ruckSack RuckSack, err error) {
	lineLength := len(line)
	if lineLength%2 != 0 || lineLength == 0 {
		return RuckSack{}, errors.New(fmt.Sprintf("%s can't be split equally into two containers", line))
	}

	ruckSack = RuckSack{line[0 : lineLength/2], line[lineLength/2 : lineLength]}
	return
}

func parseFile(filePath string) (results []RuckSack, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []RuckSack{}, err
	}

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		ruckSack, err := ruckSackFromLine(fileScanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}

		results = append(results, ruckSack)
	}

	return results, nil
}

func firstStar(ruckSacks []RuckSack) {
	priority := 0
	for _, ruckSack := range ruckSacks {
		misplacedItem, err := ruckSack.MisplacedItem()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// fmt.Fprintln(os.Stdout, "Rucksack", ruckSack, " has a misplaced ", string(misplacedItem), MisplacedItemPriority(misplacedItem))
		priority += ItemPriority(misplacedItem)
	}
	fmt.Println("First Star: ", priority)
}

func secondStar(ruckSacks []RuckSack) {
	var groups [][]RuckSack
	for index := 0; index < len(ruckSacks); index += 3 {
		if index+3 > len(ruckSacks) {
			fmt.Println("Couldn't split all ruckSacks into groups of 3")
			return
		}
		groups = append(groups, ruckSacks[index:index+3])
	}

	priority := 0
	for _, group := range groups {
		badge := FindBadge(group)
		badgePriority := ItemPriority(badge)
		//		fmt.Fprintln(os.Stdout, "Group: ", group, " badge: ", string(badge), "priority: ", badgePriority)
		priority += badgePriority
	}
	fmt.Println("Second Star: ", priority)
}

func main() {
	ruckSacks, err := parseFile("input.data")
	if err != nil {
		fmt.Println(err)
		return
	}
	firstStar(ruckSacks)
	secondStar(ruckSacks)
}
