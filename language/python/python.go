package python

import (
	"slidenerd-glot-run/cmd"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	return cmd.RunStdin(workDir, stdin, "python", files[0])
}
