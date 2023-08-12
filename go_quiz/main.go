package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type queston struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of queston, answer")

	timeLimit := flag.Int64("limit", 30, "time limit for quiz in seconds")

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

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	// <-timer.C // It will block the code as it wait mesg from channel

	correctCount := 0
	// problemLoop:
	for i, q := range questions {

		fmt.Printf("Question #%d: %s =\n", i+1, q.q)

		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {

		case <-timer.C:
			fmt.Printf("Your time limit is completed (Allocated Second : %d)\n", *timeLimit)
			fmt.Printf("You Score %d out of %d\n", correctCount, len(questions))
			// break problemLoop
			return

		case answer := <-answerCh:
			if answer == q.a {
				correctCount++
				// fmt.Printf("Correct Answer \n")
			}

		}

		// fmt.Println(questions)
	}
	fmt.Printf("You Score %d out of %d\n", correctCount, len(questions))
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
