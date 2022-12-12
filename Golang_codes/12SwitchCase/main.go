package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is one, you can open the game ")
		fallthrough //this is used , when 1 is encountered it will also print the case 2 also
	case 2:
		fmt.Println("Dice value is 2, you can move two spot")
	case 3:
		fmt.Println("Dice value is 3, you can move Three spot")
	case 4:
		fmt.Println("Dice value is 4, you can move four spot")
	case 5:
		fmt.Println("Dice value is 5, you can move five spot")
	case 6:
		fmt.Println("Dice value is 6, you can move six spot")
	default:
		fmt.Println("Value is not possible ")

	}
}
