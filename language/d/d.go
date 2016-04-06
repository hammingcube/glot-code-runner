package d

import (
	"github.com/maddyonline/glot-code-runner/cmd"
	"github.com/maddyonline/glot-code-runner/util"
	"path/filepath"
)

func Run(files []string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	binName := "a.out"

	sourceFiles := util.FilterByExtension(files, "d")
	args := append([]string{"dmd", "-of" + binName}, sourceFiles...)
	stdout, stderr, err := cmd.Run(workDir, args...)
	if err != nil {
		return stdout, stderr, err
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.Run(workDir, binPath)
}
