package main

import (
	"fmt"
	"go-by-example/arrays"
	"go-by-example/closures"
	"go-by-example/helloworld"
	"go-by-example/if-else"
	"go-by-example/looping"
	maps "go-by-example/map"
	"go-by-example/recursion"
	"go-by-example/slices"
	"go-by-example/switch"
	"go-by-example/values"
	"go-by-example/variables"
	"go-by-example/variadic"
)

func main() {

	fmt.Println("Main function in main.go")
	fmt.Println("Printing the Variables...")
	variables.Variables()
	fmt.Println()
	fmt.Println("Printing the Hello World...")
	helloworld.Helloworld()
	fmt.Println()
	fmt.Println("Printing the Values...")
	values.Values()
	fmt.Println()
	fmt.Println("Printing the For Loop examples...")
	forloop.Forloop()
	fmt.Println("Printing the If Else conditions here...")
	ifelse.Ifelse()
	fmt.Println("Printing the switch case conditions here")
	switchcase.SwitchExample()
	fmt.Println("Printing the Array examples here...")
	arrays.Arrays()
	fmt.Println("Printing the Slice examples here...")
	slices.Slices()
	fmt.Println("Printing the Map examples here...")
	maps.Map()
	fmt.Println("Using custom sum function with variadic.")
	variadic.MyCustomSum(1,2,3,4,5,6)
	closures.Closures()
	fmt.Println("Recursion function in go")
	recursion.Fab(10)


}
