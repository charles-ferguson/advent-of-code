package main

import (
	"bufio"
	"fmt"
	"os"
)

type Play interface {
	beats() Play
	beatenBy() Play
	value() int
}

type Rock struct{}

func (r Rock) beats() Play {
	return Scissors{}
}

func (r Rock) beatenBy() Play {
	return Paper{}
}

func (r Rock) value() int {
	return 1
}

type Paper struct{}

func (p Paper) beats() Play {
	return Rock{}
}

func (p Paper) beatenBy() Play {
	return Scissors{}
}

func (p Paper) value() int {
	return 2
}

type Scissors struct{}

func (s Scissors) beats() Play {
	return Paper{}
}

func (s Scissors) beatenBy() Play {
	return Rock{}
}

func (s Scissors) value() int {
	return 3
}

var myPlayMap = map[rune]Play{
	'X': Rock{},
	'Y': Paper{},
	'Z': Scissors{},
}

var oponentPlayMap = map[rune]Play{
	'A': Rock{},
	'B': Paper{},
	'C': Scissors{},
}

type Round struct {
	firstChar  rune
	secondChar rune
}

type Outcome interface {
	value() int
}

type Win struct{}

func (w Win) value() int {
	return 6
}

type Lose struct{}

func (l Lose) value() int {
	return 0
}

type Draw struct{}

func (d Draw) value() int {
	return 3
}

var charToOutcomeMap = map[rune]Outcome{
	'X': Lose{},
	'Y': Draw{},
	'Z': Win{},
}

func (r Round) firstStarScore() int {
	myPlay := myPlayMap[r.firstChar]
	oponentPlay := oponentPlayMap[r.secondChar]
	switch oponentPlay {
	case myPlay.beats():
		return Win{}.value() + myPlay.value()
	case myPlay.beatenBy():
		return Lose{}.value() + myPlay.value()
	default:
		return Draw{}.value() + myPlay.value()
	}
}

func (r Round) secondStarScore() int {
	oponentPlay := oponentPlayMap[r.secondChar]
	desiredOutcome := charToOutcomeMap[r.firstChar]

	switch desiredOutcome {
	case Win{}:
		return oponentPlay.beatenBy().value() + desiredOutcome.value()
	case Lose{}:
		return oponentPlay.beats().value() + desiredOutcome.value()
	case Draw{}:
		return oponentPlay.value() + desiredOutcome.value()
	}
	return 0
}

func roundFromLine(line string) Round {
	charLine := []rune(line)
	oponentChar := charLine[0]
	myChar := charLine[2]
	return Round{myChar, oponentChar}
}

func parseInputFile(filePath string) []Round {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	var rounds []Round
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		round := roundFromLine(fileScanner.Text())
		rounds = append(rounds, round)
	}
	return rounds
}

func firstStar(rounds []Round) {
	score := 0
	for _, round := range rounds {
		score = score + round.firstStarScore()
	}
	fmt.Fprintln(os.Stdout, "First Star: ", score)
}

func secondStar(rounds []Round) {
	score := 0
	for _, round := range rounds {
		score = score + round.secondStarScore()
	}
	fmt.Fprintln(os.Stdout, "Second Star: ", score)
}

func main() {
	rounds := parseInputFile("input.data")
	firstStar(rounds)
	secondStar(rounds)
}
