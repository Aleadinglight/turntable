package scanner

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func ScanFiles(dirToScan string, extension string) (files []string, err error) {
	err = filepath.Walk(dirToScan, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file has the corresponding extension
		if path.Ext(filePath) == extension {
			files = append(files, filePath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func ReadJSONMetadata(metadataFiles []string) ([]SongInfo, error) {
	var songInfo []SongInfo

	for _, file := range metadataFiles {
		// Open the file
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		// Read the JSON metadata
		var metadata SongMetadata
		err = json.NewDecoder(f).Decode(&metadata)
		if err != nil {
			return nil, err
		}

		// Create SongInfo from SongMetadata
		info := SongInfo{
			SongMetadata: metadata,
			FilePath:     file,
		}

		songInfo = append(songInfo, info)
	}

	return songInfo, nil
}

func GetAllSongs(dirToScan string) ([]SongInfo, error) {
	metadataFiles, err := ScanFiles(dirToScan, ".json")
	if err != nil {
		return nil, fmt.Errorf("fail to scan for metadata files in %s: %v", dirToScan, err)
	}

	songInfo, err := ReadJSONMetadata(metadataFiles)
	if err != nil {
		return nil, fmt.Errorf("fail to read metadata files: %v", err)
	}

	return songInfo, nil
}
