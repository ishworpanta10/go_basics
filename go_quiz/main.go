package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type queston struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of queston, answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s \n", *csvFilename))
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided csv file")
	}

	questions := parseLines(lines)

	correctCount := 0
	for i, q := range questions {
		fmt.Printf("Question #%d: %s =\n", i+1, q.q)

		var answer string
		fmt.Scanf("%s\n", &answer)

		if q.a == answer {
			correctCount++
			// fmt.Printf("Correct Answer \n")
		}

		fmt.Printf("You Score %d out of %d\n", correctCount, len(questions))
	}

	// fmt.Println(questions)
}

func parseLines(lines [][]string) []queston {
	ret := make([]queston, len(lines))
	for i, line := range lines {
		ret[i] = queston{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
