// main.go
package main

import (
	"fmt"
	"os"

	"github.com/aleadinglight/turntable/config"
	"github.com/aleadinglight/turntable/downloader"
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

func main() {
	rootCmd.AddCommand(cmdScan)
	rootCmd.AddCommand(cmdDownload)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
