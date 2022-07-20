package main

import "fmt"

func EvenOrOdd(number *int) string {
	if *number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {
	for {
		fmt.Print("Enter a number: ")
		var number int
		fmt.Scanf("%d\n", &number)
		fmt.Printf("%d is %s\n", number, EvenOrOdd(&number))
	}
}
