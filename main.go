package main

import (
	"log"
	"net/http"

	"sandbox/pkg/controllers"
)

const (
	PORT = ":5000"
)

// TODO: add loader when sending network requests

func main() {
	mux := http.NewServeMux()
	{
		mux.HandleFunc("/", controllers.ListAllTodosHandler)
		mux.HandleFunc("/add-todo", controllers.AddTodoHandler)
		mux.HandleFunc("/toggle", controllers.ToggleTodoStatusHandler)
	}

	log.Printf("running on port %s\n", PORT)
	http.ListenAndServe(PORT, mux)
}
