package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	for _, letter := range strings.Split(strings.TrimSpace(str), "") {
		switch letter {
		case "a", "e", "i", "o", "u":
			count++
		default:
			continue
		}
	}
	return count
}

func main() {
	for {
		var str string
		fmt.Print("Enter your string: ")
		fmt.Scanf("%s\n", &str)
		if str == "exit!" {
			fmt.Println("Goodbye :)")
			return
		} else {
			fmt.Printf("%s contains %d vowels\n", str, GetCount(str))
		}
	}
}
