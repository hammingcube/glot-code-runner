package assembly

import (
	"github.com/maddyonline/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	objName := "a.o"
	binName := "a.out"

	stdout, stderr, err := cmd.Run(workDir, "nasm", "-f", "elf64", "-o", objName, files[0])
	if err != nil {
		return stdout, stderr, err
	}

	stdout, stderr, err = cmd.Run(workDir, "ld", "-o", binName, objName)
	if err != nil {
		return stdout, stderr, err
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.Run(workDir, binPath)
}
