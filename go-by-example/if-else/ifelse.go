package ifelse

import "fmt"

func Ifelse() {
	condition := true
	// ! in a go, the else condition must be on the same line as the closing brace of the if block
	if condition {
		fmt.Println("Condition is true")
	} else {
		fmt.Println("Condition is false")
	}

	// if-else if-else ladder
	num := 10
	if num < 0 {
		fmt.Println("Number is negative")
	} else if num == 0 {
		fmt.Println("Number is zero")
	} else {
		fmt.Println("Number is positive")
	}

	// Nested if statements
	score := 85
	if score >= 0 && score <= 100 {
		if score >= 90 {
			fmt.Println("Grade: A")
		} else if score >= 80 {
			fmt.Println("Grade: B")
		} else if score >= 70 {
			fmt.Println("Grade: C")
		} else if score >= 60 {
			fmt.Println("Grade: D")
		} else {
			fmt.Println("Grade: F")
		}
	} else {
		fmt.Println("Invalid score")
	}
	
}