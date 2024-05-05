package player

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/aleadinglight/turntable/config"
)

var (
	cmd         *exec.Cmd
	lock        sync.Mutex
	pidFilePath = config.PidFilePath
)

func PlayMP3(mp3 string) error {
	lock.Lock()
	defer lock.Unlock()

	if isSongPlaying() {
		return fmt.Errorf("another song is currently playing")
	}

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

	// Write PID to a file
	err = os.WriteFile(pidFilePath, []byte(fmt.Sprintf("%d", cmd.Process.Pid)), 0644)
	if err != nil {
		return fmt.Errorf("failed to write PID file: %w", err)
	}

	return nil
}

func Stop() error {
	lock.Lock()
	defer lock.Unlock()

	if !isSongPlaying() {
		return fmt.Errorf("no song is currently playing")
	}

	// Read PID from file
	pidData, err := os.ReadFile(pidFilePath)
	if err != nil {
		return fmt.Errorf("failed to read PID file: %w", err)
	}

	pid, err := strconv.Atoi(string(pidData))
	if err != nil {
		return fmt.Errorf("invalid PID: %w", err)
	}

	if !isMPG123Process(pid) {
		return fmt.Errorf("no mpg123 process with PID: %d", pid)
	}

	// Send signal to the process with the given PID
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to find process: %w", err)
	}

	err = process.Signal(os.Interrupt)
	if err != nil {
		return fmt.Errorf("failed to stop MP3: %w", err)
	}

	// Remove the PID file
	if err := os.Remove(pidFilePath); err != nil {
		return fmt.Errorf("failed to remove PID file: %w", err)
	}

	return nil
}

func isMPG123Process(pid int) bool {
	cmdPath := fmt.Sprintf("/proc/%d/cmdline", pid)
	cmdline, err := os.ReadFile(cmdPath)
	if err != nil {
		// The process might have terminated
		return false
	}

	// Check if the cmdline contains the name of the mpg123 binary
	return strings.Contains(string(cmdline), "mpg123")
}

func isSongPlaying() bool {
	if _, err := os.Stat(pidFilePath); err == nil {
		return true // PID file exists, song is likely playing
	}
	return false // No PID file, no song is playing
}
