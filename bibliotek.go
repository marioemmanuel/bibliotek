package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
)

// Global variable to store the file path
var rootPath string

func init() {
	// Define the command line flag
	// The flag is -path, with a default value and a short description
	flag.StringVar(&rootPath, "path", "", "Path to the folder to be exposed")

	// Parse the flags
	flag.Parse()

	// Check if the path is provided
	if rootPath == "" {
		fmt.Println("You must provide a file root path using the -path flag.")
		os.Exit(1)
	}
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
    fs := http.FileServer(http.Dir("static"))
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
