package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a CSV file for starting the quiz")
	flag.Parse()
	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open CSV file: %s\n", *csvFilename)
		os.Exit(1)
	}
	_ = file
}
