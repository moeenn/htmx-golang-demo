package main

import (
	"log"
	"net/http"
	"sandbox/pkg/controllers"
)

const (
	PORT = ":3000"
)

func main() {
	mux := http.NewServeMux()

	// register all request handlers
	mux.HandleFunc("/", controllers.ListAllTodosHandler)
	mux.HandleFunc("/add-todo", controllers.AddTodoHandler)
	mux.HandleFunc("/toggle", controllers.ToggleTodoStatusHandler)

	// serve static files
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public", fs))

	log.Printf("running on port %s\n", PORT)
	http.ListenAndServe(PORT, mux)
}
