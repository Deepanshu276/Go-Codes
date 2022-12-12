package main

import "fmt"

func main() {
	Languages := make(map[int64]string)

	Languages[1] = "javascript"
	Languages[2] = "Ruby"
	Languages[3] = "Python"

	fmt.Println("List of all Laguages: ", Languages)

	fmt.Println("JS shorts for :", Languages[1])

	delete(Languages, 1)

	//fmt.Println(deletedElement)

	for key, value := range Languages {
		fmt.Printf("the key of %v is = %v\n", value, key)
	}
}
