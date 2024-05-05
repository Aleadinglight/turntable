package config

// MusicDir is the directory for reading and writing data
var MusicDir = "~/Music"

func Update(path string) {
	MusicDir = path
}
