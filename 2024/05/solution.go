package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const trace = true

type PageOrderingRule struct {
	precedingPage int
	followingPage int
}

func (rule *PageOrderingRule) updateIsCompliant(update []int) bool {
	seenFollowingPage := false
	for _, page := range update {
		if page == rule.followingPage {
			seenFollowingPage = true
		}
		if page == rule.precedingPage && seenFollowingPage {
			return false
		}
	}
	return true
}

type Update []int

func (update *Update) middlePage() int {
	if len(*update)%2 == 0 {
		return (*update)[len(*update)/2-1]
	}
	return (*update)[len(*update)/2]
}

func (update *Update) copy() Update {
	copy := make(Update, len(*update))
	for i := 0; i < len(*update); i++ {
		copy[i] = (*update)[i]
	}
	return copy
}

func (update *Update) IsCompliant(rules []PageOrderingRule) bool {
	for _, rule := range rules {
		if !rule.updateIsCompliant(*update) {
			return false
		}
	}
	return true
}

func (update *Update) reorder(rules []PageOrderingRule) Update {
	pagesToPlace := make(Update, len(*update))
	for i := 0; i < len(*update); i++ {
		pagesToPlace[i] = (*update)[i]
	}

	for len(pagesToPlace) > 0 {
		for i := 0; i < len(pagesToPlace); i++ {
			for _, rule := range rules {
				if rule.updateIsCompliant(pagesToPlace) {
					break
				}
			}
		}
	}
	return *update
}

func parseUpdate(line string) Update {
	var update Update
	for _, stringInt := range strings.Split(line, ",") {
		intInt, _ := strconv.Atoi(stringInt)
		update = append(update, intInt)
	}
	return update
}

func parsePageOrderingRule(line string) PageOrderingRule {
	var rule PageOrderingRule
	fmt.Sscanf(line, "%d|%d", &rule.precedingPage, &rule.followingPage)
	return rule
}

func parseInput(filepath string) ([]PageOrderingRule, []Update) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil
	}

	rules := make([]PageOrderingRule, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		} else {
			rules = append(rules, parsePageOrderingRule(line))
		}
	}

	updates := make([]Update, 0)
	for scanner.Scan() {
		line := scanner.Text()
		updates = append(updates, parseUpdate(line))
	}
	return rules, updates
}

func firstStar(rules []PageOrderingRule, updates []Update) int {
	sum := 0
	updateCompliant := true
	for _, update := range updates {
		updateCompliant = true
		for _, rule := range rules {
			if !rule.updateIsCompliant(update) {
				if trace {
					fmt.Print("Update not compliant: ")
					for _, page := range update {
						colorReset := "\033[0m"
						colorRed := "\033[31m"
						if page == rule.precedingPage || page == rule.followingPage {
							fmt.Print(colorRed, page, colorReset, ",")
						} else {
							fmt.Print(page, ",")
						}
					}
				}
				fmt.Println("Rule: ", rule)
				updateCompliant = false
				break
			}
		}
		if updateCompliant {
			fmt.Println("Update compliant: ", update)
			if trace {
				fmt.Println("Update compliant: ", update, "middle page: ", update.middlePage())
			}
			sum += update.middlePage()
		}
	}
	return sum
}

func secondStar() {
	sum := 0
	for _, update := range updates {
		for _, rule := range rules {
			if rule.updateIsCompliant(update) {
				sum += update.middlePage()
				break
			}
		}
	}
}
func main() {
	rules, updates := parseInput("input.data")
	fmt.Println("Rules: ", len(rules))
	fmt.Println("Updates: ", len(updates))
	fmt.Println("First star: ", firstStar(rules, updates))
}
