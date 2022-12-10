package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza store")

	fmt.Println("Thanks for your valuable rating")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	println("Thanks for giving us", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to rating ", numrating+1)
	}

}
