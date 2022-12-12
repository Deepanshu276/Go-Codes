/*Time Complexity ---------> Undefined
As f(n)=O(g(n))
f(n)==infinity , thus it will not map n to real number,hence f(n) is not a real valued function

*/
// Space Complexity----------> O(1) (The variable declared go out of scope everytime, they get destroyed , hence space complexity remains constant)
package main

//Entry Point of the Program and tells the Go compiler that the package should compile as an executable program instead of a shared library.

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	NAME := bufio.NewReader(os.Stdin) //Taking name from the User

	AGE := bufio.NewReader(os.Stdin) //Taking age from the User

	COUNTRY := bufio.NewReader(os.Stdin) // Taking Country of the User

	name, _ := NAME.ReadString('\n') //Reading the name given by the User till the new line . '\n ' acts as a delimiter
	age, _ := AGE.ReadString('\n')
	country, _ := COUNTRY.ReadString('\n') //Reading the country given by the User till the new line . '\n ' acts as a delimiter

	for {
		fmt.Printf("Your name is %v Your age is %v Years old and Your nationality is  %v", name, age, country)
	}

}
