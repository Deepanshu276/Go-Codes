package main

import "fmt"

//jwtToken:=2000 This type of decleration only works inside the function

const LoginPage string = "hdkjdksjdk" // The first character in Loginpage is Capital which means it is public in nature

func main() {
	var name string = "Deepanshu is my name"
	fmt.Println(name)
	fmt.Printf("The Type of this name is %T \n", name)

	var age int = 909202920
	fmt.Println(age)
	fmt.Printf("the type of this age is %T \n", age)

	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("The Boolean state of this isTrue is : %T \n", isTrue)

	//implicit Type
	var website = "redhat"
	fmt.Println(website)

	//No var style

	numberofuser := 30000
	fmt.Println(numberofuser)

	fmt.Println(LoginPage)
	fmt.Printf("The variable is of type : %T \n", LoginPage)

}
