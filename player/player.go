package player

import (
	"fmt"
	"os"
	"os/exec"
)

func PlayMP3(mp3 string) error {
	// Check if mpg123 command is available
	_, err := exec.LookPath("mpg123")
	if err != nil {
		return fmt.Errorf("mpg123 command not found: %w", err)
	}

	cmd := exec.Command("mpg123", mp3)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to play MP3: %w", err)
	}

	return nil
}
