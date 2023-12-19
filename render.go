package main

import (
	"regexp"
	"sort"
	"strings"
    "fmt"
    "io/ioutil"
    "path/filepath"
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/parser"
)

// RenderMarkdownFile renders a markdown file to HTML
func RenderMarkdownFile(mdFilePath, rootPath string) (string, error) {
    // Compute the full path
    fullPath := filepath.Join(rootPath, mdFilePath)

    // Read the markdown file content
    mdContent, err := ioutil.ReadFile(fullPath)
    if err != nil {
        return "", err
    }

    // Create a new markdown parser with common extensions
    extensions := parser.CommonExtensions
    mdParser := parser.NewWithExtensions(extensions)
	
    // Convert markdown content to HTML
    htmlContent := markdown.ToHTML(mdContent, mdParser, nil)
    return string(htmlContent), nil
}

// RenderTree recursively builds HTML content for the file tree
func RenderTree(node FileNode, basePath string, cabinetLevel bool) string {
    var htmlContent string

    // Sort children alphabetically
	// They shall be already sortened, but this is to ensure that they are
    sort.Slice(node.Children, func(i, j int) bool {
        return node.Children[i].Name < node.Children[j].Name
    })

    // Directly render the children of the root node
    if basePath == "" { // Check if it's the root
	
        // Start the root <ul> tag
        htmlContent += "<ul>"
        for _, child := range node.Children {
            htmlContent += RenderTree(child, child.Name, true)
        }
        // Close the root <ul> tag
        htmlContent += "</ul>"
    } else {
		// Check and strip the "DDD_" prefix from the name for display
		var displayName string
	    re := regexp.MustCompile(`^(\d{3})_`)
		if re.MatchString(node.Name) {
			splitName := strings.SplitN(node.Name, "_", 2)
			if len(splitName) == 2 {
				displayName = splitName[1]
			} else {
				displayName = node.Name
			}
		} else {
			displayName = node.Name
		}

		// Remove the ".md" extension from the display name
    	if strings.HasSuffix(displayName, ".md") {
        	displayName = strings.TrimSuffix(displayName, ".md")
    	}

        // For folders and files under the root
        if node.IsFolder {
            folderPath := basePath 

			if cabinetLevel {
            	// Caret, cabinet icon (close/open) and Folder Name constitute the clickable folder item
           		htmlContent += fmt.Sprintf("<li><span class='folder-item'><span class='caret'></span><span class='folderclose'>🖿</span><span class='folderopen'>🗁</span> %s</span>", displayName)
			} else {
           		htmlContent += fmt.Sprintf("<li><span class='folder-item'><span class='caret'></span><span class='folderclose'>🖿</span><span class='folderopen'>🗁</span> %s</span>", displayName)
			}
            htmlContent += "<ul class='nested'>"

            for _, child := range node.Children {
                // Recursively render child folders and files
                htmlContent += RenderTree(child, folderPath, false)
            }
            htmlContent += "</ul></li>"
        } else {

		    // Prepend a Unicode character for document
		    //documentIcon := "<span class='symbol'>🗏 </span>"
		    //documentIcon := "<span class='symbol'>🗎 </span>"
		    documentIcon := "<span class='symbol'>🗍 </span>"
		    displayName = documentIcon + displayName

            // File as a clickable link
			htmlContent += fmt.Sprintf("<li><a href='javascript:void(0);' content-path='/%s'>%s</a></li>", node.Path, displayName)	
        }
    }
    return htmlContent
}
