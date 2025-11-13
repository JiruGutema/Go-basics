package closures

import "fmt"

// Closures are functions that reference variables from outside their body.
// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
// this is the same as anonymous function but with the ability to access variables from the outer scope.

func intIncrement() func() int {
	counter := 0

	// the structure of the closure is a variable assiged to a unnamed function. varibale name := func() return type { }

	increment := func() int {
		counter++
		return counter
	}

	return increment
}

func Closures() {
	fmt.Println("This is the Closures example function.")
	increment := intIncrement()

	// each time we call increment, it will remember the value of counter from the previous call.

	fmt.Println("Counter: ", increment()) // Counter: 1
	fmt.Println("Counter: ", increment())
	fmt.Println("Counter: ", increment())
	fmt.Println("Counter: ", increment())
	fmt.Println("Counter: ", increment()) // Counter: 5

	newIncrement := intIncrement() // create a new closure instance
	fmt.Println("New Counter: ", newIncrement()) // New Counter: 1
	fmt.Println("New Counter: ", newIncrement())
	fmt.Println("New Counter: ", newIncrement())

}

// closures are very useful as we don't nest named functions in go unlike other programming languages. Instead, we can use closures to encapsulate state and behavior together.



