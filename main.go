package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"sandbox/pkg/controllers"
	"sandbox/pkg/templates"
)

const (
	ADDRESS = "0.0.0.0:3000"
)

//go:embed public
var publicFiles embed.FS

//go:embed views
var viewsFiles embed.FS

func main() {
	mux := http.NewServeMux()

	/* register all request handlers */
	mux.HandleFunc("/", controllers.ListAllUsersHandler)
	mux.HandleFunc("/add-user", controllers.AddUserHandler)
	mux.HandleFunc("/toggle", controllers.ToggleUserStatusHandler)
	mux.HandleFunc("/remove", controllers.UserRemoveHandle)

	/* serve static files and templates */
	publicFS := http.FS(publicFiles)
	templates.Initialize(viewsFiles)
	mux.Handle("/public/", http.FileServer(publicFS))

	log.Printf("running on %s\n", ADDRESS)
	if err := http.ListenAndServe(ADDRESS, mux); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start web server: %s\n", err.Error())
		os.Exit(1)
	}
}
