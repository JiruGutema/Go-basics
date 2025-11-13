package variadic

import "fmt"
// variadic function is a function that takes a variable number of arguments.
// we can define a variadic function by using ... before the type of the last parameter.
// inside the function, the variadic parameter is treated as a slice of that type.

func MyCustomSum(numbers ...int) int {
	fmt.Println("numbers to be summed: ", numbers)
	fmt.Println("inside my customo variadic function. ")
	total := 0

	for _, num := range numbers {
		total += num
	}
	fmt.Println("total: ", total)
	return total
}