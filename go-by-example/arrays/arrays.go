package arrays

import (
	"fmt"
)

func Arrays() {
	// we declare an array in a go, by specifying the size and type of elements
	var arr [5]int // an array of 5 integers initialized to zero values
	fmt.Println("This is an empty array: ", arr)

	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	
	fmt.Println("This is an array after assigning values: ", arr)
	// we can also declare and initialize an array in a single line
	arr2 := [3]string{"Go", "is", "awesome"}
	fmt.Println("This is another array: ", arr2)

	// we can also use ellipsis to let go determine the size of the array
	arr3 := [...]float64{3.14, 1.618, 2.718}
	fmt.Println("This is an array with ellipsis: ", arr3)

	// we can iterate over the array using for loop
	fmt.Println("Iterating over arr2:")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Element at index %d: %s\n", i, arr2[i])
	}

	// we can also use range to iterate over the array
	fmt.Println("Iterating over arr3 using range:")
	for index, value := range arr3 {
		fmt.Printf("Element at index %d: %f\n", index, value)
	}	

	twoD := [2][3]int{
		{1, 2,3},
		{4,5,6},
	}

	fmt.Println("This is the 2d Array: ", twoD)

	var matrix [3][4]int

	n := 3
	m := 4

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			matrix[r][c] = r * c
		}
	}

	fmt.Println("This is the matrix after populating values: ", matrix)
}