package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "This is my first user input in golang"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our course")
	// Cooma ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank for giving star", input)

	fmt.Printf("Thank yor for providing the rating %T \n", input)

}
