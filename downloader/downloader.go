package downloader

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/aleadinglight/turntable/config"
	"github.com/google/uuid"
)

const (
	DownloaderApp   = config.DownloaderApp
	AudioProcessApp = config.AudioProcessApp
)

// checks if DownloaderApp and ffmpeg are installed on the system
func checkDependencies() error {
	// Check if DownloaderApp is installed
	_, err := exec.LookPath(DownloaderApp)
	if err != nil {
		return fmt.Errorf("%s is not installed", DownloaderApp)
	}

	// Check if ffmpeg is installed
	_, err = exec.LookPath(AudioProcessApp)
	if err != nil {
		return fmt.Errorf("%s is not installed", AudioProcessApp)
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

// downloads the audio from a YouTube video as an MP3 file and returns the file name
func DownloadMP3(youtubeVideoLink string, saveDir string) (dir string, fileName string, err error) {
	dir = saveDir
	if dir == "" {
		dir = config.MusicDir
	}

	fileName = uuid.New().String()

	err = runCommand(DownloaderApp, "-x", "--audio-format", "mp3", "-P", dir, "-o", fileName, "--write-info-json", youtubeVideoLink)
	if err != nil {
		err = fmt.Errorf("failed to download MP3: %v", err)
	}

	return dir, fileName, err
}
