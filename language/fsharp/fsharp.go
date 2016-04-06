package fsharp

import (
	"github.com/maddyonline/glot-code-runner/cmd"
	"github.com/maddyonline/glot-code-runner/util"
	"path/filepath"
)

func Run(files []string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	binName := "a.exe"

	sourceFiles := reverse(util.FilterByExtension(files, "fs"))
	args := append([]string{"fsharpc", "--out:" + binName}, sourceFiles...)
	stdout, stderr, err := cmd.Run(workDir, args...)
	if err != nil {
		return stdout, stderr, err
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.Run(workDir, "mono", binPath)
}

func reverse(files []string) []string {
	reversed := make([]string, 0, len(files))

	for i := len(files) - 1; i >= 0; i-- {
		reversed = append(reversed, files[i])
	}

	return reversed
}
