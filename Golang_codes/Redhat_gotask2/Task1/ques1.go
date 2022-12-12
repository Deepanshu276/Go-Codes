package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Unix is used to is used to yield “t” as a Unix time that is the number of seconds passed
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100) + 1

	if randomNumber > 50 {
		fmt.Print("It's closer to 100\n")

	} else {
		fmt.Print("It's closer to 0\n")
	}

}
