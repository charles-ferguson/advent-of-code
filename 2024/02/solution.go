package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Report []int

func (report Report) createDampenedReport(index int) Report {
	dampenedReport := make(Report, len(report)-1)
	for i := 0; i < len(report)-1; i++ {
		if i < index {
			dampenedReport[i] = report[i]
		} else {
			dampenedReport[i] = report[i+1]
		}
	}
	return dampenedReport
}

func (report Report) isMonotonic() (bool, int) {
	isIncreasing := report[0] < report[len(report)-1]
	for i := 1; i < len(report); i++ {
		if isIncreasing && report[i] < report[i-1] {
			return false, i
		}
		if !isIncreasing && report[i] > report[i-1] {
			return false, i
		}
	}
	return true, -1
}

func (report Report) allIncrementsAreSafe() (bool, int) {
	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i] - report[i-1]))
		if diff < 1 || diff > 3 {
			return false, i
		}
	}
	return true, -1
}

func parseReports(filePath string) []Report {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	reports := make([]Report, 0)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		report := make(Report, 0)
		line := fileScanner.Text()
		words := strings.Fields(line)

		for _, word := range words {
			intValue, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println("Invalid input")
				return make([]Report, 0)
			}
			report = append(report, intValue)
		}
		reports = append(reports, report)
	}
	return reports
}

func countSafeReports(reports []Report, allowance int) int {
	count := 0
	for _, report := range reports {
		exceptions := 0

		for exceptions <= allowance {
			reportIsMonotonic, index := report.isMonotonic()
			if reportIsMonotonic {
				break
			}

			exceptions++
			report = report.createDampenedReport(index)
		}

		for exceptions <= allowance {
			reportAllIncrementsAreSafe, index := report.allIncrementsAreSafe()
			if reportAllIncrementsAreSafe {
				break
			}

			exceptions++
			report = report.createDampenedReport(index)
		}

		if exceptions <= allowance {
			count++
		}
	}
	return count
}

func FirstStar() {
	reports := parseReports("input.data")
	fmt.Printf("Safe Reports: %d\n", countSafeReports(reports, 0))
}

func SecondStar() {
	reports := parseReports("input.data")
	fmt.Printf("Safe Damped Reports: %d\n", countSafeReports(reports, 1))
}

func main() {
	FirstStar()
	SecondStar()
}
