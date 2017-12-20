package cpp

import (
	"slidenerd-glot-run/cmd"
	"slidenerd-glot-run/util"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	binName := "a.out"

	sourceFiles := util.FilterByExtension(files, "cpp")
	args := append([]string{"clang++", "-std=c++11", "-o", binName}, sourceFiles...)
	stdout, stderr, err := cmd.Run(workDir, args...)
	if err != nil {
		return stdout, stderr, err
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.RunStdin(workDir, stdin, binPath)
}
