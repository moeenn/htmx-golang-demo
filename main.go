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
	mux.HandleFunc("/", controllers.ListAllUsersHandler)
	mux.HandleFunc("/add-user", controllers.AddUserHandler)
	mux.HandleFunc("/toggle", controllers.ToggleUserStatusHandler)
	mux.HandleFunc("/remove", controllers.UserRemoveHandle)

	// serve static files
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public", fs))

	log.Printf("running on port %s\n", PORT)
	http.ListenAndServe(PORT, mux)
}
