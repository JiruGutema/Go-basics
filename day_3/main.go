package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go is running")
	operators() // calling the operators function
}

func operators() {
	// arithmetic operators
	a := 10
	b := 20
	sum := a + b
	fmt.Println("Sum:", sum)
	diff := b - a
	fmt.Println("Difference:", diff)
	product := a * b
	fmt.Println("Product:", product)
	quotient := b / a
	fmt.Println("Quotient:", quotient)
	remainder := b % a
	fmt.Println("Remainder:", remainder)
	// relational operators
	isEqual := a == b
	fmt.Println("Is Equal:", isEqual)
	isNotEqual := a != b
	fmt.Println("Is Not Equal:", isNotEqual)
	isGreater := a > b
	fmt.Println("Is Greater:", isGreater)
	isLess := a < b
	fmt.Println("Is Less:", isLess)
	isGreaterOrEqual := a >= b
	fmt.Println("Is Greater Or Equal:", isGreaterOrEqual)
	isLessOrEqual := a <= b
	fmt.Println("Is Less Or Equal:", isLessOrEqual)
	// logical operators
	isTrue := true
	isFalse := false
	andResult := isTrue && isFalse
	fmt.Println("Logical AND:", andResult)
	orResult := isTrue || isFalse
	fmt.Println("Logical OR:", orResult)
}
