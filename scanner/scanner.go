package scanner

import (
	"os"
	"path"
	"path/filepath"
)

func ScanMP3Files(dirToScan string) ([]string, error) {
	var mp3Files []string

	err := filepath.Walk(dirToScan, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file has an mp3 extension
		if path.Ext(filePath) == ".json" {
			mp3Files = append(mp3Files, filePath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return mp3Files, nil
}
