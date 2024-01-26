package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sandbox/pkg/controllers"
)

const (
	ADDRESS = "0.0.0.0:3000"
)

func main() {
	mux := http.NewServeMux()

	/* register all request handlers */
	mux.HandleFunc("/", controllers.ListAllUsersHandler)
	mux.HandleFunc("/add-user", controllers.AddUserHandler)
	mux.HandleFunc("/toggle", controllers.ToggleUserStatusHandler)
	mux.HandleFunc("/remove", controllers.UserRemoveHandle)

	/* serve static files */
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public", fs))

	log.Printf("running on %s\n", ADDRESS)
	if err := http.ListenAndServe(ADDRESS, mux); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start web server: %s\n", err.Error())
		os.Exit(1)
	}
}
