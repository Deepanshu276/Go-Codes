package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is notes of the Slice Lecture")

	var fruits = []string{"Apple", "Bnana", "Cheeku"}

	fmt.Println("The list of Fruits is \n", fruits)
	fmt.Println("The Length of the Fruits \n", len(fruits))

	fruits = append(fruits, "orange") //we can only append string in the string slice

	fmt.Println("The New Fruits List is ", fruits)

	fmt.Println("The new fruits List is ", fruits[1:3])

	//Method 2

	Scores := make([]int, 4)

	Scores[0] = 323
	Scores[1] = 212
	Scores[2] = 321
	Scores[3] = 123

	fmt.Println(Scores)

	Scores = append(Scores, 3232) //We can append the score but can't add manually
	fmt.Println(Scores)

	sort.Ints(Scores)
	fmt.Println(Scores)
	fmt.Println(sort.IntsAreSorted(Scores))

	//Sorting in reverse Order
	sort.Slice(Scores, func(i, j int) bool {
		return Scores[i] > Scores[j]
	})
	for _, v := range Scores {
		fmt.Println(v)
	}

	//How to remove value from Slices based on index
	index := 2

	Scores = append(Scores[:index], Scores[index+1:]...)
	fmt.Println(Scores)

}
