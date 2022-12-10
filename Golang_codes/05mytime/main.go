package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time module class")
	//used for the present time
	present_time := time.Now()
	fmt.Println(present_time)

	//for formatting the time accorsdig to our format, this used date, time, day is by default and we have to use this only
	fmt.Println(present_time.Format("01-02-2006 15:04:05 Monday"))

	//Creating our own date

	createdDate := time.Date(2020, time.April, 10, 2, 1, 2, 1, time.UTC)

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

}
