package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to sihunta"
	fmt.Println(welcome)

	// Taking input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// this is comma ok or error ok syntax
	// In golang we don't have try catch statement. error ok or comma ok syntax will take care of it.
	input, _ := reader.ReadString('\n') // u either gonna get a input or might recieve an error
	fmt.Println("Thanks for reading", input)
}
