package main

import "fmt"

func main() {
	fmt.Print("This is the Pointers Code\n")

	//var ptr *int
	//fmt.Println("The value of the pointer is ", ptr)

	name := "Deepanshu"

	var ptr = &name

	fmt.Println("The Address of the Pointer is ", ptr)

	fmt.Println("The value of the Pointer is ", *ptr)

	*ptr = *ptr + " " + "Jha"

	fmt.Println("The New Value of the Pointer is ", *ptr)
}
