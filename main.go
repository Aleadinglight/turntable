// main.go
package main

import (
	"fmt"

	"github.com/aleadinglight/turntable/scanner"
)

func main() {
	const SoundDir = "./sounds"

	// Scan all mp3 files in the sounds directory
	mp3Files, err := scanner.ScanMP3Files(SoundDir)

	if err != nil {
		fmt.Printf("Error when scanning directory '%s'. Error: %v\n", SoundDir, err)
		return
	}

	fmt.Printf("Found %d mp3 files:\n", len(mp3Files))
	for _, mp3File := range mp3Files {
		fmt.Println(mp3File)
	}
}
