package routes

import (
	"fmt"
	"go-rest-api/controllers"
	"net/http"
)

func StudentRoutes() {
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.HandleGetStudent(w, r)
		case http.MethodPost:
			controllers.HandlePostStudent(w, r)
		case http.MethodPut:
			controllers.HandlePutStudent(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Student routes registered on /students")
}