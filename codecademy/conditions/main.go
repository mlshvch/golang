package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but rememeber, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)
	if isHeistOn && (openedVault >= 70) {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("Vault can't be opened :(")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		isHeistOn = false
		switch leftSafely {
		case 0:
			fmt.Println("Heist is failed :(")
		case 1:
			fmt.Println("The car haven't arrived")
		case 2:
			fmt.Println("Vault doors don't open from the inside")
		default:
			isHeistOn = true
			fmt.Println("Start gateway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("We grabed %v dollars!!!", amtStolen)
	}

	fmt.Println(isHeistOn)
}
