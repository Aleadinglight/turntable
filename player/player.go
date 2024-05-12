package player

import (
	"fmt"
	"sync"

	"github.com/aleadinglight/turntable/utils"
)

const (
	musicApp = "mpg123"
)

var (
	lock sync.Mutex
)

func PlayMP3(mp3 string) error {
	lock.Lock()
	defer lock.Unlock()

	// Check if mpg123 command is available
	if !utils.CheckAppExists(musicApp) {
		return fmt.Errorf("mpg123 is not installed")
	}

	// TODO Check if another song is playing

	err := utils.RunCommandAsyncWithArgs("mpg123", []string{mp3})
	if err != nil {
		return fmt.Errorf("failed to play MP3: %w", err)
	}

	return nil
}

// Stop all currently playing songs
func Stop() error {
	lock.Lock()
	defer lock.Unlock()

	// Run command to stop playing MP3
	err := utils.RunCommandWithArgs("pkill", []string{musicApp})
	if err != nil {
		return fmt.Errorf("failed to stop playing MP3: %w", err)
	}

	return nil
}
