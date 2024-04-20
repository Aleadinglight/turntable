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
	const dirToScan = "./sounds"
	var mp3Files []string

	// Walk through the directory
	err := filepath.Walk(dirToScan, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file has an mp3 extension
		if path.Ext(filePath) == ".mp3" {
			mp3Files = append(mp3Files, filePath)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error when scanning", dirToScan, ". Error:", err)
		return
	}

	// Print the overall result length and list of mp3 files
	fmt.Printf("Found %d mp3 files:\n", len(mp3Files))
	for _, mp3File := range mp3Files {
		fmt.Println(mp3File)
	}
}
