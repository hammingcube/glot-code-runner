package cmd

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
)

func Run(workDir string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = workDir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if stdin := getStdin(workDir); stdin != nil {
		cmd.Stdin = stdin
		defer stdin.Close()
	}
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func RunBash(workDir, command string) (string, string, error) {
	return Run(workDir, "bash", "-c", command)
}

func getStdin(workDir string) *os.File {
	const STDIN_FILE_NAME = "_stdin_"
	stdinFile := filepath.Join(workDir, STDIN_FILE_NAME)
	if _, err := os.Stat(stdinFile); err == nil {
		if stdin, err := os.Open(stdinFile); err == nil {
			return stdin
		}
	}
	return nil
}
