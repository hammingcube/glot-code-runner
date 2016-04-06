package idris

import (
	"path/filepath"
	"github.com/maddyonline/glot-code-runner/cmd"
)

func Run(files []string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	binName := "a.out"

	stdout, stderr, err := cmd.Run(workDir, "idris", "-o", binName, files[0])
	if err != nil {
		return stdout, stderr, err
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.Run(workDir, binPath)
}
