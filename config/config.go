package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	MusicPlayerApp  = "mpg123"
	DownloaderApp   = "yt-dlp"
	AudioProcessApp = "ffmpeg"
)

var (
	// MusicDir is the directory for reading and writing data
	MusicDir    string
	PidFilePath string
)

func init() {
	// Get the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	// Construct the full path to the Music directory
	MusicDir = filepath.Join(homeDir, "Music")
	PidFilePath = filepath.Join(MusicDir, "temp", "turntable_mpg123.pid")
}

func UpdateMusicDir(newDir string) {
	MusicDir = newDir
	PidFilePath = filepath.Join(MusicDir, "temp", "turntable_mpg123.pid")
}
