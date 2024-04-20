package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	// Define the directory to scan
	// You can replace "./" with any directory path
	dirToScan := "./sounds"

	// Walk through the directory
	filepath.Walk(dirToScan, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error:", err)
			return err
		}

		// Check if the file has an mp3 extension
		if path.Ext(filePath) == ".mp3" {
			fmt.Println(filePath)
		}
		return nil
	})
}
