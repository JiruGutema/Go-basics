package functions

import "fmt"
// functions are a block of code that performs a specific task.
// to declare a function we use the func keyword followed by the function name and parentheses.
// inside the parentheses we can define parameters if needed.
// the function body is enclosed in curly braces.

// unlike python, we can't have functions inside functions in go.

func Functions(){
	fmt.Println("This is the Functions example function.")

	a, b := vals()

	fmt.Println("printing from the multi variable", a, b)
}

func Add(a int, b int) int {
	return a + b
}

func Swap(x, y string) (string, string) {
	return y, x
}

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// naked return -> means returning the named return values
	return
}

// for the multiple return values, use the type as a tuple in the function signature. 
func vals()(int, int){
	return 3, 7
}
