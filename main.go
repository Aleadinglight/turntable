// main.go
package main

import (
	"flag"
	"fmt"

	"github.com/aleadinglight/turntable/scanner"
)

func main() {
	// Define a string flag with a default value and a usage description
	soundDir := flag.String("d", "./sounds", "the directory to scan for mp3 files")

	// Parse the flags
	flag.Parse()

	// Now that flag.Parse() has been called, we can use the value of soundDir
	SoundDir := *soundDir

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
