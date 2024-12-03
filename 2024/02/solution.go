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

func countSafeReports(reports []Report) int {
	count := 0
	for _, report := range reports {
		if !reportIsMonotonic(report) {
			continue
		}

		if !allIncrementsAreSafe(report) {
			continue
		}
		count++
	}
	return count
}

func dampedReportSafe(report Report) bool {
	for i := 0; i < len(report); i++ {
		dampenedReport := make(Report, len(report)-1)
		for j := 0; j < len(report)-1; j++ {
			if i > j {
				dampenedReport[j] = report[j]
			} else if i <= j {
				dampenedReport[j] = report[j+1]
			}
		}
		if reportIsMonotonic(dampenedReport) && allIncrementsAreSafe(dampenedReport) {
			return true
		}
	}
	return false
}

func countSafeReportsWithDamper(reports []Report) int {
	count := 0
	for _, report := range reports {
		if dampedReportSafe(report) {
			count++
		}
	}
	return count
}

func allIncrementsAreSafe(report Report) bool {

	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i] - report[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}
func reportIsMonotonic(report Report) bool {
	isIncreasing := report[0] < report[len(report)-1]
	for i := 1; i < len(report); i++ {
		if isIncreasing && report[i] < report[i-1] {
			return false
		}
		if !isIncreasing && report[i] > report[i-1] {
			return false
		}
	}
	return true
}

func FirstStar() {
	reports := parseReports("input.data")
	fmt.Printf("Safe Reports: %d\n", countSafeReports(reports))
}

func SecondStar() {
	reports := parseReports("input.data")
	fmt.Printf("Safe Damped Reports: %d\n", countSafeReportsWithDamper(reports))
}

func main() {
	FirstStar()
	SecondStar()
}
