// backend/main.go
package main

import (
	"fmt"
	"groupie-tracker/backend/handlers"
	"html/template"
	"net/http"
	"path/filepath"
)

// main function, entry point of the program
func main() {
	// Serve static files (styles and images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/styles"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./frontend/images"))))

	// Define routes for handling requests
	http.HandleFunc("/", handleNotFound)
	http.HandleFunc("/500", handle500)

	// Goroutine to fetch artists in the background
	go func() {
		_, err := handlers.GetArtists()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Artists Fetched")
	}()

	// Set the port for the server to listen on
	port := "3000"
	println("Server listening on port http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

// Function to render HTML templates
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Filepaths for template and layout files
	tmpl = filepath.Join("frontend", tmpl+".html")
	layout := filepath.Join("frontend", "layout.html")

	// Parse template files
	t, err := template.ParseFiles(layout, tmpl)
	if err != nil {
		// Handle template parsing error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Execute template with data
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		// Handle template execution error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Handler for 404 Not Found errors
func handleNotFound(w http.ResponseWriter, r *http.Request) {
	// Check if the requested path is the root path
	if r.URL.Path == "/" {
		// Get artists and their relations
		combinedData, err := handlers.GetArtistsWithRelations()
		if err != nil {
			fmt.Println("Error:", err)
			http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
			return
		}
		// Render the index page with the retrieved data
		renderTemplate(w, "index", combinedData)
	} else {
		// Render the 404 page
		data := struct{}{}
		renderTemplate(w, "404", data)
	}
}

// Handler for 500 Internal Server Error
func handle500(w http.ResponseWriter, r *http.Request) {
	// Render the 500 page
	data := struct{}{}
	renderTemplate(w, "500", data)
}
