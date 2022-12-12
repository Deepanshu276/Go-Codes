// struct has no Inheritance and Parent child concept
package main

import "fmt"

func main() {
	Deepanshu := user{"Deepanshu", "abc@gmail.com", true, 34}

	fmt.Println(Deepanshu)
	fmt.Printf("Deepanshu's Details are :%v\n", Deepanshu)
	fmt.Printf("Name is %v and email is %v", Deepanshu.Name, Deepanshu.Email)

}

type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
