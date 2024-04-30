package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// checks if yt-dlp and ffmpeg are installed on the system
func checkDependencies() error {
	// Check if yt-dlp is installed
	_, err := exec.LookPath("yt-dlp")
	if err != nil {
		return fmt.Errorf("yt-dlp is not installed")
	}

	// Check if ffmpeg is installed
	_, err = exec.LookPath("ffmpeg")
	if err != nil {
		return fmt.Errorf("ffmpeg is not installed")
	}

	return nil
}

// runs a command and streams its output to the console
func runCommand(command string, args ...string) error {
	err := checkDependencies()
	if err != nil {
		fmt.Printf("Dependency check failed: %v\n", err)
		return fmt.Errorf("missing dependencies  %v", err)
	}

	// Create a new command
	cmd := exec.Command(command, args...)
	// Create a pipe for the command's stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %v", err)
	}

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start command: %v", err)
	}

	go func() {
		// Stream the output of the command
		io.Copy(os.Stdout, stdout)
	}()

	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("command failed: %v", err)
	}

	return nil
}

// downloads the audio from a YouTube video as an MP3 file
func DownloadMP3(youtubeVideoLink string) error {
	err := runCommand("yt-dlp", "-x", "--audio-format", "mp3", youtubeVideoLink)
	if err != nil {
		return fmt.Errorf("failed to download MP3: %v", err)
	}

	return nil
}
