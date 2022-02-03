package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers class")

	var ptr *int // this is how you initialize pointers
	fmt.Println("The value of pointer is ", ptr)

	myNum := 23
	var ptr2 = &myNum // passing the memory address
	fmt.Println("the address where myNum is stored", ptr2)
	fmt.Println("the value by derefrencing the pointer is", *ptr2) //De-refrencing the pointer

	*ptr2 = *ptr2 + 2
	fmt.Println("the chandged value is ", *ptr2) //We are perform the changes in the actual memory location
}
