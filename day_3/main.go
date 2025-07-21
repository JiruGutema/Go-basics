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


	day := 8
  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  default:
    fmt.Println("Not a weekday")
  }



   switch day {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }
}

