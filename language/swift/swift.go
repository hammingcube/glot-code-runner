package swift

import (
	"path/filepath"
	"github.com/maddyonline/glot-code-runner/util"
	"github.com/maddyonline/glot-code-runner/cmd"
)

func Run(files []string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	binName := "a.out"

	sourceFiles := util.FilterByExtension(files, "swift")
	args := append([]string{"swiftc", "-o", binName}, sourceFiles...)
	stdout, stderr, err := cmd.Run(workDir, args...)
	if err != nil {
		return stdout, stderr, err
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.Run(workDir, binPath)
}
