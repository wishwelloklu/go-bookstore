package routes

import (
	"github.com/gorilla/mux"
	// "github.com/akhil/go-bookstore/pkg/controllers"
	"github.com/wishwelloklu/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(route *mux.Route) {
	route.HandlerFunc("/book/", controllers.CreateBook).Methods("POST")
	route.HandlerFunc("/book/", controllers.GetBook).Methods("GET")
	route.HandlerFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	route.HandlerFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	route.HandlerFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
