package forloop

import "fmt"

func Forloop() {
	// Basic for loop in Go
	// ! 1. Initialization; Condition; Post statement

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration: ", i)
	}

	//! 2. Condition only. this is one is similar to the while loop
	var j int = 0
	for j < 5 {
		fmt.Println("Iterating ", j, "times")
		j++
	}

	//! 3. Infinite loop
	// Uncomment the following lines to see infinite loop in action

	// for {
	// 	fmt.Println("This is an infinite loop")
	// }

	// To stop the infinite loop, you can use a break statement or a return statement
	condition := true
	for condition {
		fmt.Println("This loop will run until condition is false")
		condition = false
	}

	// Using break statement to exit infinite loop
	condition = true
	count := 0
	for {
		fmt.Println("Breaking out of infinite loop")
		if !condition {
			break
		}
		count++
		if count >= 1 {
			condition = false
		}
	}
	//! Using range keyword to iterate over a collection
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Println("Index: ", index, "Value: ", value)
	}

	//! Ignoring index or value using blank identifier _
	for _, value := range numbers {
		fmt.Println("Value only: ", value)
	}
	// Ignoring value using blank identifier _
	for index := range numbers { //! go will ignore the value by default
		fmt.Println("Index only implicit: ", index)
	}

	for index, _ := range numbers { //! go will ignore the value by default
		fmt.Println("Index only explicit: ", index)
	}

	nums := 10

	for num := range nums { //! iterating over the range starting from the 0 all the way to nums -1
		fmt.Println("num is ", num)
	}

}

func Whileloop() {
	// While loop in Go using for loop

}
