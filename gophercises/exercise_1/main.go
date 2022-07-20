package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a CSV file for starting the quiz")
	timeLimit := flag.Int("time", 30, "quiz time limit")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse provided CSV file :(")
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	problems := parseLines(lines)
	correctAnswers := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("\n---Time is over---")
			showResult(&correctAnswers, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.answer {
				correctAnswers++
			}
		}
	}

	showResult(&correctAnswers, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func showResult(result *int, totalScore int) {
	fmt.Printf("You scored %d out of %d.\n", *result, totalScore)
}
