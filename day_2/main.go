package main
import ("fmt")

func main(){
	go_arrays()
	go_slices()
}

func go_arrays(){

	var ppl = [4]string{"Jiru", "Gutema", "Biftu", "Dasi"} // using var name = [length]type{values}
	fmt.Println("My Arrays: ", ppl)

	var animals = [...]string{"dog", "cat","cow"} // here length is infered
	fmt.Println(animals)

	nums := [3]int{1,2,3} // using := [length]type{values}
	fmt.Println(nums)

	nums_1 := [...]int{1,2,3,4,5} // here length is infered
	fmt.Println(nums_1)

	arr1 := [5]int{} //not initialized
	arr2 := [5]int{1,2} //partially initialized
	arr3 := [5]int{1,2,3,4,5} //fully initialized
	

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	arr4 := [5]int{1:10,2:40} // initializes the specifi elements of the array. first and third. 

  	fmt.Println(arr4)

	fmt.Println(len(arr4))

}

func go_slices(){ // unlike arrays, slices can be grow and shrink in sizes
	// declare the slices
	var nums = []string{}
	nums = append(nums, "Gutema")
	nums = append(nums, "Jiru")
	fmt.Println(nums)

}