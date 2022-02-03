package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time story of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	// Now this where golang gets crazy
	// If u want to format the string u can use the format fxn but u have to provide a standard input to get the result
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // u can skip anypart u don't want but be careful of the typecase
	createDate := time.Date(2011, time.December, 12, 8, 32, 67, 22, time.UTC)
	fmt.Println(createDate)
	// To format again we have to do the same
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}
