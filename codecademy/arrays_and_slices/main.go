package main

import "fmt"

func printFirstLastArray(array [4]int) {
	fmt.Println("First", array[0])
	fmt.Println("Last", array[3])
}

func changeFirst(slice []int, value int) {
	if len(slice) > 0 {
		slice[0] = value
	}
}

func main() {
	myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	fmt.Println(myTutors)
	theirHours := [4]int{20, 10, 30, 15}
	theirHoursSlice := theirHours[:]
	changeFirst(theirHoursSlice, 25)
	fmt.Println(theirHoursSlice)
	printFirstLastArray(theirHours)
}
