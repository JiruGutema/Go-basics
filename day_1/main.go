package main // declaring the package name for the file. main is the package for this package

import (
	"fmt"    
)  // importing a fmt package to use in this code

func main() { // function declaretion name. main is the name of this function
	
	variable() // calling a variables func
	dataTypes() // calling a datatypes func
}

// a := 10 // causes error. declaring the it outside causes an error.
func variable(){
	var student string = "Jiru Gutema" //declaring a variable named student of type string  and initilizing. we can omit one of them or all of them
	fmt.Println(student) // using a Print fun from the package fmt to print out the result

	dog := "bobby" // here we are declaring a variable called dog. type is infered at runtime. value must be given in this case
	fmt.Println(dog) // printing the value of the dog. type infered at the run time.
	// ---------
	a := 10 // since the variable is declared inside the block, it works fine.
	fmt.Println(a)
	// -----------
	// multiple variable declaration
	var d,b,c string = "a","b","c" //we can declare the same type if we specify the type
	fmt.Println(d,b,c)
	// ------
	// var e,f,g int = 10, "jiru", 10 // this is error. type should be the same
	// println(e,f,g)
	// -------
	// to do the above case, we can use the := declaration method to declare multiple var with diffferent types
	g,h := "Jiru", 10
	fmt.Println(g,h)
	// ------ using var()
	var(
		j = "string"
		k int = 10
	)
	fmt.Println(j,k)
	// ------ using a constant
	const PI = 3.14
	fmt.Println(PI)
	// -------- using formatted print
	fmt.Printf("%d,%v,%s,%t", 10, PI, j, true)

}

func dataTypes(){
	var a bool = true     // Boolean
	var b int = 5         // Integer
	var c float32 = 3.14  // Floating point number
	var d string = "Hi!"  // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
}