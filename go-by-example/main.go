package main

import (
	"fmt"
	"go-by-example/values"
	"go-by-example/variables"
	"go-by-example/helloworld"
	"go-by-example/looping"
	"go-by-example/if-else"
	"go-by-example/switch"
	"go-by-example/arrays"
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

}