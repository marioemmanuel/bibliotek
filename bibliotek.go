// PROJECT: bibliotek makefile
// AUTHOR: MARIO EMMANUEL
// DATE: 2023/DEC/19
// github.com/marioemmanuel/bibliotek

package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
)

// Global variable to store the file path
var basePath string
var rootPath string
var staticPath string

func init() {
	// Define the command line flag
	// The flag is -path, with a default value and a short description
	flag.StringVar(&basePath, "path", "", "Path to the folder to be exposed")

	// Parse the flags
	flag.Parse()

	// Check if the path is provided
	if basePath == "" {
		fmt.Println("You must provide a file base path using the -path flag.")
		fmt.Println("This path will contain two folders: static and content")
		fmt.Println("static will contain the HTML, CSS and JS files")
		fmt.Println("content will contain your provided structure of folders, files and markdown files")
		os.Exit(1)
	}

	// Populate content and static paths
	rootPath = basePath + "/content"
	staticPath = basePath + "/static"
}

func main() {

	// Use the filePath in your application
	fmt.Printf("Exposing files in folder: %s\n", rootPath)

    fileTree, err := BuildFileTree(rootPath)
    if err != nil {
        fmt.Println("Error building file tree:", err)
        os.Exit(1)
    }

    // Serve static files
    fs := http.FileServer(http.Dir(staticPath))
    http.Handle("/", fs)

    http.HandleFunc("/content/", renderMarkdownHandler)
	http.HandleFunc("/tree", renderTreeHandler)
    
	fmt.Printf("\nfileTree:\n---\n%+v\n---\n", fileTree)
   
	// Start Web server 
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }

}
