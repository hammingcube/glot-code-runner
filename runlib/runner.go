package runlib

import (
	"encoding/json"
	"fmt"
	"github.com/maddyonline/glot-code-runner/cmd"
	"github.com/maddyonline/glot-code-runner/language"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Payload struct {
	Language string          `json:"language"`
	Files    []*InMemoryFile `json:"files"`
	Command  string          `json:"command"`
}

type InMemoryFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}

func Run(input_stdin io.Reader, output_stdout io.Writer) {
	payload := &Payload{}
	err := json.NewDecoder(input_stdin).Decode(payload)

	if err != nil {
		exitF("Failed to parse input json (%s)\n", err.Error())
	}

	// Ensure that we have at least one file
	if len(payload.Files) == 0 {
		exitF("No files given\n")
	}

	// Check if we support given language
	if !language.IsSupported(payload.Language) {
		exitF("Language '%s' is not supported\n", payload.Language)
	}

	// Write files to disk
	filepaths, err := writeFiles(payload.Files)
	if err != nil {
		exitF("Failed to write file to disk (%s)", err.Error())
	}

	var stdout, stderr string

	// Execute the given command or run the code with
	// the language runner if no command is given
	if payload.Command == "" {
		stdout, stderr, err = language.Run(payload.Language, filepaths)
	} else {
		workDir := filepath.Dir(filepaths[0])
		stdout, stderr, err = cmd.RunBash(workDir, payload.Command)
	}
	printResult2(stdout, stderr, err, output_stdout)
}

// Writes files to disk, returns list of absolute filepaths
func writeFiles(files []*InMemoryFile) ([]string, error) {
	// Create temp dir
	tmpPath, err := ioutil.TempDir("", "")
	if err != nil {
		return nil, err
	}

	paths := make([]string, len(files), len(files))
	for i, file := range files {
		path, err := writeFile(tmpPath, file)
		if err != nil {
			return nil, err
		}

		paths[i] = path

	}
	return paths, nil
}

// Writes a single file to disk
func writeFile(basePath string, file *InMemoryFile) (string, error) {
	// Get absolute path to file inside basePath
	absPath := filepath.Join(basePath, file.Name)

	// Create all parent dirs
	err := os.MkdirAll(filepath.Dir(absPath), 0775)
	if err != nil {
		return "", err
	}

	// Write file to disk
	err = ioutil.WriteFile(absPath, []byte(file.Content), 0664)
	if err != nil {
		return "", err
	}

	// Return absolute path to file
	return absPath, nil
}

func exitF(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func printResult(stdout, stderr string, err error) {
	result := &Result{
		Stdout: stdout,
		Stderr: stderr,
		Error:  errToStr(err),
	}
	json.NewEncoder(os.Stdout).Encode(result)
}

func printResult2(stdout, stderr string, err error, output_stdout io.Writer) {
	result := &Result{
		Stdout: stdout,
		Stderr: stderr,
		Error:  errToStr(err),
	}
	json.NewEncoder(output_stdout).Encode(result)
}

func errToStr(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}
