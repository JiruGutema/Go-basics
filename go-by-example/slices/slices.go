package slices

import "fmt"

// slices are a dynamic array in go. letting us having dynamic length of the elements.

func Slices() {
	fmt.Println("This is the Slices example function.")
	var slice = []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice before appending values: ", slice)
	slice = append(slice, 7, 8, 9)
	slice = append(slice, 10, 11, 12)
	fmt.Println("slice after appending values: ", slice)
	fmt.Println()
	fmt.Println("Iterating through the slice:")
	for num := range slice {
		fmt.Println(num)

	}

	var a = []string{"Go", "Python", "Java", "C++", "JavaScript"}
	fmt.Println("Original slice:", a)

	// Slicing the slice from index 1 to 3 (excluding index 3)
	b := a[1:3]
	fmt.Println("Sliced slice (1 to 3):", b)

	a = append(a, b...)
	fmt.Println("Slice after appending sliced elements:", a)

	var number100 []int

	for i := 0; i < 101; i++ {
		number100 = append(number100, i)
	}
	fmt.Println("numbers in range 100: ", number100)

	//using make to inilize with the default value

	var usingmake = make([]int, 3)
	fmt.Println("declaring slice using make", usingmake)

	// getting the length of the slice

	fmt.Println("length of the slice using len(): ", len(usingmake))
}
