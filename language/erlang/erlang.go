package erlang

import (
	"path/filepath"
	"github.com/maddyonline/glot-code-runner/cmd"
)

func Run(files []string) (string, string, error) {
	workDir := filepath.Dir(files[0])

	// Compile all files except the first
	for _, file := range files[1:] {
		stdout, stderr, err := cmd.Run(workDir, "erlc", file)
		if err != nil {
			return stdout, stderr, err
		}
	}

	// Run first file with escript
	return cmd.Run(workDir, "escript", files[0])
}
