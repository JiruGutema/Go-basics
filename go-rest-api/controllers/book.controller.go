package controllers
import (
	"fmt"
	"net/http"
)

var books = [...]string{"Book1", "Book2", "Book3"}

func HandleGetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a GET request")
	fmt.Fprintf(w, "Books available: %v\nPath: %s", books, r.URL.Path)
}

func HandlePostBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a POST request")
	fmt.Fprintf(w, "This is a response for POST request at path %s", r.URL.Path)
}

func HandlePutBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received a PUT request")
	fmt.Fprintf(w, "This is a response for PUT request at path %s", r.URL.Path)
}

