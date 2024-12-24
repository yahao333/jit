package git

import (
	"os/exec"
)

func Push() error {
	cmd := exec.Command("git", "push")
	return cmd.Run()
}
