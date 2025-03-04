package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	_ = csvFileName
	file, err := os.Open(*csvFileName)
	defer file.Close()
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	// reader for the csv file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	parseLines(lines)
	// fmt.Println(lines)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines))

	for i, line := range lines {
		res[i] = problem{
			q: line[0],
			a: line[1],
		}

		fmt.Println(res[i])
	}

	return res

}
