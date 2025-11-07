package variables

import "fmt"

func Variables() {
	// var declares 1 or more variables
	var name string = "Jiru"
	var age int = 23
	var isStudent bool = true

	var isFemale, isMarrried, hasDegree bool = false, false, true

	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Is Student: ", isStudent)
	fmt.Println("Is Female: ", isFemale)
	fmt.Println("Is Married: ", isMarrried)
	fmt.Println("Has Degree: ", hasDegree)

	// The := syntax is shorthand for declaring and initializing a variable
	city := "New York"
	country := "USA"

	fmt.Println("City: ", city)
	fmt.Println("Country: ", country)

	city, country = "Los Angeles", "USA"
	fmt.Println("Updated City: ", city)
	fmt.Println("Updated Country: ", country)

	// Short variable declarations can only be used inside functions
	//! we can't use short variable declaration outside functions
	// language := "Python" // This will cause an error
	// type while declaring variables using shorthand syntax outside functions

	//! Ignoring values using blank identifier _. Here we are ignoring the first return value
	_, arg2 := func() (int, int) {
		return 10, 20
	}()

	fmt.Println("value in interest Value: ", arg2)

	// Constants
	const pi = 3.14
	const euler = 2.71

	fmt.Println("Value of Pi: ", pi)
	fmt.Println("Value of Euler's Number: ", euler)

	// we use const keyword to declare constants that cannot be changed

}
