package main

import "fmt"

func main() {
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "The Login Count is 10"

	}
	fmt.Println(result)

	//Method 2
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")

	} else {
		fmt.Println("Num is not less than 10")
	}
}
