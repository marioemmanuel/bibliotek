package main

import (
    "os"
	"fmt"
    "net/http"
    "strings"
)

func renderMarkdownHandler(w http.ResponseWriter, r *http.Request) {
    // Extracting the markdown file path from the URL
    mdFilePath := strings.TrimPrefix(r.URL.Path, "/content/")

    // Render the markdown file
    htmlContent, err := RenderMarkdownFile(mdFilePath, rootPath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Write the HTML response
    fmt.Fprint(w, htmlContent)
}

func renderTreeHandler(w http.ResponseWriter, r *http.Request) {
   
   	// Get fileTree 
	fileTree, err := BuildFileTree(rootPath)
    if err != nil {
        fmt.Println("Error building file tree:", err)
        os.Exit(1)
    }
    
	// Use the fileTree as needed, e.g., print it for debugging
    fmt.Printf("\nfileTree:\n---\n%+v\n---\n", fileTree)

	// Convert the global fileTree to HTML
    htmlContent := RenderTree(fileTree, "", true)

    // Write the HTML response
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, htmlContent)
}
