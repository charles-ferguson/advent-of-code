package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type List []int

func findDifference(list1 List, list2 List) int {
	slices.Sort(list1)
	slices.Sort(list2)

	absDiff := float64(0)
	for index, value := range list1 {
		absDiff += math.Abs(float64(value) - float64(list2[index]))
	}
	return int(absDiff)
}

func similarityScore(list1 List, list2 List) int {
	score := 0
	slices.Sort(list1)
	slices.Sort(list2)
	for _, value := range list1 {
		count := 0
		for _, value2 := range list2 {
			if value == value2 {
				count++
			} else if value < value2 {
				break
			}
		}

		score += count * value
	}
	return score
}

func parseLists(filePath string) []List {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	lists := make([]List, 2)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		words := strings.Fields(line)

		if len(words) != 2 {
			fmt.Println("Invalid input")
			return make([]List, 0)
		}

		for index, word := range words {
			intValue, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println("Invalid input")
				return make([]List, 0)
			}
			lists[index] = append(lists[index], intValue)
		}
	}
	return lists
}

func FirstStar() {
	lists := parseLists("input.data")
	fmt.Printf("difference: %d\n", findDifference(lists[0], lists[1]))
}

func SecondStar() {
	lists := parseLists("input.data")
	fmt.Printf("similarity score: %d\n", similarityScore(lists[0], lists[1]))
}

func main() {
	FirstStar()
	SecondStar()
}
