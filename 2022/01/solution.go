package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Elf struct {
	snack_calories []int
}

func (elf Elf) TotalCalories() int {
	sum := 0
	for _, calories := range elf.snack_calories {
		sum += calories
	}
	return sum
}

func NewElf(snack_calories []int) Elf {
	dupe_calories := snack_calories
	return Elf{dupe_calories}
}

func parseElves(filePath string) []Elf {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var snack_calories []int
	var elves []Elf
	for fileScanner.Scan() {
		if fileScanner.Text() == "" {
			elves = append(elves, NewElf(snack_calories))
			snack_calories = nil
		} else {
			calories, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Println(err)
				return []Elf{}
			}

			snack_calories = append(snack_calories, calories)
		}
	}
	return elves
}

func FirstStar() {
	elves := parseElves("input.data")
	max_calories := 0
	for _, elf := range elves {
		if elf.TotalCalories() > max_calories {
			max_calories = elf.TotalCalories()
		}
	}
	fmt.Println(max_calories)
}

func SecondStar() {
	elves := parseElves("input.data")
	var top_three = [...]int{0, 0, 0}
	var calories int
	for _, elf := range elves {
		calories = elf.TotalCalories()
		if calories > top_three[0] {
			top_three[2] = top_three[1]
			top_three[1] = top_three[0]
			top_three[0] = calories
		} else if calories > top_three[1] {
			top_three[2] = top_three[1]
			top_three[1] = calories
		} else if calories > top_three[2] {
			top_three[2] = calories
		}

		fmt.Fprintln(os.Stdout, "Second Star:", top_three)
	}

	fmt.Fprintln(os.Stdout, "Second Star: ", top_three, top_three[0]+top_three[1]+top_three[2])
}

func main() {
	FirstStar()
	SecondStar()
}
