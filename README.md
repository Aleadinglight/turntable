# Turntable

Turntable is a fast and flexible music player and downloader built using Go. It allows users to play and download MP3 songs from Youtube, manage music files locally, and view song metadata in a well-formatted table.

## Features

- **Play MP3 songs** directly from your filesystem.
- **Download songs** from provided Youtube URLs into a specified directory.
- **Scan directories** for all downloaded songs and metadata information.
- **List all audio files** in a directory.
- **Stop playback** of the current song using OS signals.

## Installation

To install Turntable, you need to have Go installed on your machine. Follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/aleadinglight/turntable.git
   cd turntable
   ```

2. Build the project:
   ```bash
   go build -o turntable main.go
   ```

3. Run the executable:
   ```bash
   ./turntable
   ```

## Usage

Turntable is built on the Cobra CLI framework, providing a simple command-line interface.

### Playing a Song

To play a song:
```bash
./turntable play path/to/your/song.mp3
```

### Downloading a Song

To download a song from a URL:
```bash
./turntable download [url]
```
Replace `[url]` with the actual URL of the song.

### Listing Songs in a Directory

To list all song files in the configured directory:
```bash
./turntable list
```
This command is particularly useful for further processing, such as scripting or integration with other software tools, as it outputs a simple list of song file paths.

### Scanning a Directory for Song Metadata

To scan a directory and display detailed metadata for all songs:
```bash
./turntable menu
```
This command provides a comprehensive overview, ideal for users who need to view extensive details about each song, including titles, durations, and file names. It displays the information in a formatted table, making it easy to read and understand.

### Stopping Playback

To stop playing the current song:
```bash
./turntable stop
```

## Configuration

Modify the configuration settings in `config/config.go` to set the default music directory and other preferences.

## Contributing

Contributions to Turntable are welcome! Please fork the repository and submit pull requests with your features or fixes.

## License

Distributed under the MIT License. See `LICENSE` for more information.
