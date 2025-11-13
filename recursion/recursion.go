package recursion

import "fmt"

// Recursion is a programming technique where a function calls itself in order to solve a problem.
// A recursive function typically has two main components: a base case and a recursive case.
// The base case is the condition under which the function stops calling itself, preventing infinite recursion.
// The recursive case is where the function calls itself with modified arguments, gradually approaching the base case.

func Fab(n int) int {
	fmt.Println("fab called with n =", n)
	if n <= 1 {
		return n
	}

	fmt.Println("fab returning:", Fab(n - 1) + Fab(n - 2))
	return Fab(n - 1) + Fab(n - 2)
}