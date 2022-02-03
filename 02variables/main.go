package main

import "fmt"

const LoginToken = "C255" // This is how you declare a constant
// Here Capital 'L' of LoginToken has a meaning. It means that LoginToken variable is public

func main() {
	var username string = "harry"
	fmt.Println(username)
	fmt.Printf("Variable is of the type %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of the type %T \n", isLoggedin)

	var smallint uint16 = 45681
	fmt.Println(smallint)
	fmt.Printf("Variable is of the type %T \n", smallint)

	var smallfloat = 23.81238906
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of the type %T \n", smallfloat)

	// Default values and some alias
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of the type %T \n", anotherVariable)

	// implicit type
	var website = "instagram.com" // Here we do not need to mention the type. lexer does it automatically for us.
	fmt.Println(website)          // same as python

	// no var style
	noOfUsers := 3000      // We mainly use this style in the project work
	fmt.Println(noOfUsers) // We can olny do this inside a method. Declaring global variables like this is not allowed.

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of the type %T \n", LoginToken)

}
