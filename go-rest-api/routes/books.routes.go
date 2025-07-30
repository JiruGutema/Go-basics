package routes

import (
	"fmt"
	"go-rest-api/controllers"
	"net/http"
)
func BooksRoutes(){
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.HandleGetBooks(w, r)
		case http.MethodPost:
			controllers.HandlePostBooks(w, r)
		case http.MethodPut:
			controllers.HandlePutBooks(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Books routes registered on /books")
}