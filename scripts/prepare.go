package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	distPath := filepath.Join("frontend", "dist")
	
	// Create directory
	err := os.MkdirAll(distPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Create placeholder index.html
	indexPath := filepath.Join(distPath, "index.html")
	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		err = os.WriteFile(indexPath, []byte("<!DOCTYPE html><html><body>Placeholder</body></html>"), 0644)
		if err != nil {
			fmt.Printf("Error creating placeholder: %v\n", err)
			os.Exit(1)
		}
	}
	
	fmt.Println("Project environment prepared.")
}
