// PROJECT: bibliotek makefile
// AUTHOR: MARIO EMMANUEL
// DATE: 2023/DEC/19
// github.com/marioemmanuel/bibliotek

package main

import (
    "io/ioutil"
    "path/filepath"
    "strings"
)

// FileNode struct and BuildFileTree function
type FileNode struct {
    Name     string      // Name of the file or folder
    Path     string      // Relative path to the file or folder
    IsFolder bool        // Flag to identify if it's a folder
    Children []FileNode  // Child nodes (for folders)
    Assets   []string    // List of asset files (specific to .md files)
}

func BuildFileTree(rootPath string) (FileNode, error) {
    return buildTree(rootPath, "")
}

func buildTree(basePath, currentPath string) (FileNode, error) {
    fullPath := filepath.Join(basePath, currentPath)
    fileInfo, err := ioutil.ReadDir(fullPath)
    if err != nil {
        return FileNode{}, err
    }

    var node FileNode
    node.Path = currentPath
    node.IsFolder = true

    // Extract the folder name from the currentPath
    if currentPath != "" {
        _, node.Name = filepath.Split(currentPath)
    }

    for _, file := range fileInfo {
        relativePath := filepath.Join(currentPath, file.Name())

        if file.IsDir() {
            childNode, err := buildTree(basePath, relativePath)
            if err != nil {
                return FileNode{}, err
            }
            node.Children = append(node.Children, childNode)
        } else if strings.HasSuffix(file.Name(), ".md") {
            mdNode := FileNode{
                Name:     file.Name(),
                Path:     relativePath,
                IsFolder: false,
            }

            assetsFolderPath := file.Name() + ".files"
            assetsPath := filepath.Join(fullPath, assetsFolderPath)
            if assets, err := ioutil.ReadDir(assetsPath); err == nil {
                for _, asset := range assets {
                    assetRelativePath := filepath.Join(relativePath, assetsFolderPath, asset.Name())
                    mdNode.Assets = append(mdNode.Assets, assetRelativePath)
                }
            }

            node.Children = append(node.Children, mdNode)
        }
    }

    return node, nil
}
