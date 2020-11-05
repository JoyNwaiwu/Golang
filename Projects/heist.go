package main

import (
	"fmt"
	"math/rand"
	"time"
)

var heistOn bool = true

func main() {
	rand.Seed(time.Now().UnixNano())
	escapingGuards()
	vaultAccess()
	successfulOperation()
	fmt.Println("\nThe Heist is:", heistOn)
}

func escapingGuards() {
	eludedGuards := rand.Intn(100)
	if eludedGuards >= 50 {
		fmt.Println("\nYou managed to make it past the guards. Good job, but this is just the first step.")
	} else {
	heistOn = false
		fmt.Println("\nOops! Plan a better disguise next time, failed heist!")
		return
	}
}

func vaultAccess() {
	openedVault := rand.Intn(100)
	if heistOn == true && openedVault >= 70 {
		fmt.Println("\nThe vault is open!")
		fmt.Println("Grab and Go!!!")
	} else {
	heistOn = false
		fmt.Println("Sorry bro, this vault can't be opened")
	}
}

func successfulOperation() {
	leftSafely := rand.Intn(4)

	if heistOn == true {
		switch leftSafely {
		case 0:
			heistOn = false
			fmt.Print("\nLooks like you tripped an alarm... run!!!")
		case 1:
			heistOn = false
			fmt.Print("\nOh no! we got trapped inside the vault, it doesn't open from inside...failed mission!!!")
		case 2:
			heistOn = false
			fmt.Print("This is the Police freeze!")
		default:
			fmt.Println("\nWe've got the money...start the truck!")
			amountStolen := 10000 + rand.Intn(1000000)
			fmt.Print("We are $",amountStolen,"+ richer....hahaahah!!!")
		}
	}

}







