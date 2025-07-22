package main 
import ("fmt")

func main() {

	var name string
	fmt.Println("What is your name?")
	fmt.Scan(&name);
	var no_subjects int
	fmt.Println("Number of subjects")
	fmt.Scan(&no_subjects)
	
	
	type StudentMark struct {
		name string
		grade string
	}
	var marks = []StudentMark{}
	var tot int = 1

	for  i:=0; i < no_subjects; i ++{
			var mark int
			var sub_name string
			fmt.Println("Please Enter for the", i, "th subjects name")
			fmt.Scanln(&sub_name)
			fmt.Println("Please Enter for the", i, "th subjects mark")
			fmt.Scanln(&mark)

			if (mark >= 80){
				marks = append(marks, StudentMark{name: sub_name, grade: "A"})
			}else if (mark >= 70){
				marks = append(marks, StudentMark{name:sub_name, grade: "B"})
			} else if (mark >= 60){
				marks = append(marks, StudentMark{name: sub_name, grade: "C"})
			}else if (mark >= 50) {
				marks = append(marks, StudentMark{name: sub_name, grade: "D"})
			} else{
				marks = append(marks, StudentMark{name: sub_name, grade: "F"})
			}
			tot += mark

		}
	var average int = tot / no_subjects


	fmt.Println("Student name", name)
	fmt.Println("Average: ", average)

	fmt.Println("Student Subject Grade: ")

	for i:=0; i < no_subjects; i++{
		fmt.Println(marks[i].name, ":", marks[i].grade)
	}


	
	

	
}

