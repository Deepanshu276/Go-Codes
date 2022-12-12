package main

import "fmt"

func main() {
	fmt.Print("This is Array Lecture\n")

	var Fruits [4]string

	Fruits[0] = "Apple"
	Fruits[1] = "Banana"
	Fruits[3] = "Pine Apple"
	Fruits[2] = "Orange"

	fmt.Print("The Fruit List is \n", Fruits, "\n")

	//If we have don't fill the element then it will leave the space for it in the output
	//Array length will be the length initialized at the starting irrespective of the number of the lement present in the array
	var Veg [4]string

	Veg[0] = "Potato"
	Veg[1] = "Tomato"
	Veg[3] = "Onion"

	fmt.Println("The veg list is ", Veg)

	fmt.Println("The length of the veg list is ", len(Veg))

	//Method 2

	var pet_number = [4]int{1, 2, 3, 4}

	fmt.Println("The pet number is ", pet_number)

}
