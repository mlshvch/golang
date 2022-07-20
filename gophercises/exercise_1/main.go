package main

import (
	"flag"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a CSV file for starting the quiz")
	flag.Parse()
	_ = csvFilename
}
