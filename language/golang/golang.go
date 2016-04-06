package golang

import (
	"github.com/maddyonline/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	return cmd.Run(workDir, "go", "run", files[0])
}
