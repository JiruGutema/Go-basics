package main

import (
	"fmt"
	"go-rest-api/routes"
	"net/http"
)

func main() {

	routes.StudentRoutes()
	routes.BooksRoutes()

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
