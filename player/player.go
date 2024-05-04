package player

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

var (
	cmd  *exec.Cmd
	lock sync.Mutex
)

func PlayMP3(mp3 string) error {
	lock.Lock()
	defer lock.Unlock()
	// Check if mpg123 command is available
	_, err := exec.LookPath("mpg123")
	if err != nil {
		return fmt.Errorf("mpg123 command not found: %w", err)
	}

	// Check if another song is playing
	if cmd != nil && cmd.Process != nil {
		return fmt.Errorf("another song is currently playing")
	}

	cmd = exec.Command("mpg123", mp3)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to play MP3: %w", err)
	}

	return nil
}

func Stop() error {
	lock.Lock()
	defer lock.Unlock()

	if cmd != nil && cmd.Process != nil {
		// Kill the process using the PID
		err := cmd.Process.Kill()
		if err != nil {
			return fmt.Errorf("failed to stop MP3: %w", err)
		}
		cmd = nil
		return nil
	}

	return fmt.Errorf("no song is currently playing")
}
