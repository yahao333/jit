package git

import (
	"os/exec"

	"github.com/yahao333/jit/internal/errors"
)

// GetStagedDiff 获取git staged的文件差异
func GetStagedDiff() (string, error) {
	// 检查是否是git仓库
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	if err := cmd.Run(); err != nil {
		return "", errors.ErrNotGitRepo
	}

	// 获取staged文件的diff
	cmd = exec.Command("git", "diff", "--staged")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	if len(output) == 0 {
		return "", errors.ErrNoChanges
	}

	return string(output), nil
}
