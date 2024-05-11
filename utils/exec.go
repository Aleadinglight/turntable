package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(command string) error {
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}

	return nil
}

func RunCommandAsync(command string) error {
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}

	return nil
}

func RunCommandWithArgs(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}

	return nil
}

func RunCommandAsyncWithArgs(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}

	return nil
}

func CheckAppExists(appName string) bool {
	_, err := exec.LookPath(appName)
	if err != nil {
		return false
	}
	return true
}
