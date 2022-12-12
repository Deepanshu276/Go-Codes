package main

import "fmt"

func main() {
	days := []string{"sunday", "Tuesday", "Thursday", "Friday"}

	fmt.Println(days)

	//Method 1
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	//method 2

	for i := range days {
		fmt.Println(days[i])
	}

	//method 3
	for index, ele := range days {
		fmt.Printf("the index of %v is %v\n", ele, index)
	}

	//Method 4
	value := 1
	for value < 10 {
		if value == 7 {
			goto lco
		}
		if value == 3 {
			value++
			continue
		}
		fmt.Println(value)
		value++
	}
	//goto
lco:
	fmt.Println("Jumping to goto method")

	//For infinte Loop
	for {
		fmt.Printf("This is an infinte loop")
	}

}
