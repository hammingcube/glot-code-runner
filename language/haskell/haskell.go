package haskell

import (
	"path/filepath"
	"github.com/maddyonline/glot-code-runner/cmd"
)

func Run(files []string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	return cmd.Run(workDir, "runghc", files[0])
}
