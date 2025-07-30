package controllers

import ("net/http"; "fmt")

var students = [...]string{"Jiru", "Chala", "Burka"}
func HandlePutStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a PUT request")
	fmt.Fprintf(w, "This is a response for PUT request at path %s", r.URL.Path)
}
func HandleGetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a GET request")
	fmt.Fprintf(w, "Student available: %v\nPath: %s", students, r.URL.Path)
}

func HandlePostStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a POST request")
	fmt.Fprintf(w, "This is a response for POST request at path %s", r.URL.Path)
}

