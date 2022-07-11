package main

import (
	"fmt"
)

func main() {

	count := 0
	for {
		fmt.Println(count)
		if count == 10 {
			break
		}
		count++
	}

	for tally := 0; tally < 20; tally++ {
		if tally == 10 {
			fmt.Println("Skipping 10!")
			continue
		}
		fmt.Println(tally)
	}

	letters := []string{"A", "B", "C", "D"}
	for index, value := range letters {
		fmt.Println("Index:", index, "Value:", value)
	}

	addressBook := map[string]string{
		"John":   "12 Main St",
		"Janet":  "56 Pleasant St",
		"Jordan": "88 Liberty Ln",
	}
	for key, value := range addressBook {
		fmt.Println("Name:", key, "Address:", value)
	}

}
