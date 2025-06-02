package routes

import (
	"github.com/gorilla/mux"
	"github.com/lakmalniranga/todo-backend/controllers"
)

// TodoRoutes registers all todo routes
func TodoRoutes(router *mux.Router) {
	// Create a todo
	router.HandleFunc("/api/todos", controllers.CreateTodo).Methods("POST")
	
	// Get all todos
	router.HandleFunc("/api/todos", controllers.GetTodos).Methods("GET")
	
	// Get a single todo
	router.HandleFunc("/api/todos/{id}", controllers.GetTodo).Methods("GET")
	
	// Update a todo
	router.HandleFunc("/api/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	
	// Delete a todo
	router.HandleFunc("/api/todos/{id}", controllers.DeleteTodo).Methods("DELETE")
}