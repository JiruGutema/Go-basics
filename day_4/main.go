package main
import ("fmt")

func main(){
	conditions()
}

func conditions(){
	a := 3
	b := 4

	if(a > b){
		fmt.Println(a,"is greater than",b)
	} else {
		fmt.Println(b, "is greater than",a)
	}

	c := 14
	d := 14
	if c < d {
		fmt.Println("a is less than b.")
	} else if c > d {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}
}
