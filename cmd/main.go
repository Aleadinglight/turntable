// main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aleadinglight/turntable/config"
	"github.com/aleadinglight/turntable/downloader"
	"github.com/aleadinglight/turntable/player"
	"github.com/aleadinglight/turntable/scanner"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "turntable",
	Short: "turntable is a tool for playing and downloading mp3 songs from the internet.",
	Long: `A Fast and Flexible Music Player and Downloader built with
            love by aleadinglight using Go.`,
}

var cmdScan = &cobra.Command{
	Use:   "scan",
	Short: "Scan for videos",
	Long:  `Scan the directory specified for music mp3 files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Access configuration for directory
		songs, err := scanner.ScanMP3Files(config.MusicDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error scanning for songs in %s: %v\n", config.MusicDir, err)
			os.Exit(1)
		}
		fmt.Println("Found songs:", songs)
	},
}

var cmdDownload = &cobra.Command{
	Use:   "download [url]",
	Short: "Download a song from a specified URL.",
	Long:  `Download a song from the specified URL. It will be saved in the default directory.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0] // Get the URL from the arguments
		err := downloader.DownloadMP3(url, config.MusicDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error downloading video from %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("Downloaded video from:", url)
	},
}

var playCmd = &cobra.Command{
	Use:   "play [path]",
	Short: "Play an MP3 file",
	Long:  `Plays the specified MP3 file using the player package.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := player.PlayMP3(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error playing MP3 file: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Playing:", args[0])
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop playing the current song",
	Long:  `Stop playing the current song using the os signal.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		err := player.Stop()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error stopping current playing song: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Stopped", args[0])
	},
}

func main() {
	rootCmd.AddCommand(cmdScan)
	rootCmd.AddCommand(cmdDownload)
	rootCmd.AddCommand(playCmd)
	rootCmd.AddCommand(stopCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Setup signal handling to catch SIGINT and SIGTERM
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		// Clean up: Delete PID file
		os.Remove("/tmp/mpg123.pid")
		os.Exit(0)
	}()
}
