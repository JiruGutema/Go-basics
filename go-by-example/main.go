package main

import (
	"fmt"
	"go-by-example/values"
	"go-by-example/variables"
	"go-by-example/helloworld"
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

}