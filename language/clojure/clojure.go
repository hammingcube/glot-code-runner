package clojure

import (
	"github.com/maddyonline/glot-code-runner/cmd"
	"path/filepath"
)

func Run(files []string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	return cmd.Run(workDir, "java", "-cp", "/usr/share/java/clojure.jar", "clojure.main", files[0])
}
