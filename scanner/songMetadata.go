package scanner

// VideoMetadata holds only the necessary fields from the JSON
type SongMetadata struct {
	YoutubeID      string `json:"id"`
	Title          string `json:"title"`
	Duration       int    `json:"duration"`
	DurationString string `json:"duration_string"`
	Format         string `json:"format"`
	AudioChannels  int    `json:"audio_channels"`
	FileSize       int    `json:"filesize"`
}

type SongInfo struct {
	SongMetadata SongMetadata
	FullPath     string
	FileName     string
}
