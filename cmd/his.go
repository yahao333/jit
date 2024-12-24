package cmd

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var limit int

var hisCmd = &cobra.Command{
	Use:   "his",
	Short: "Display the git commit history",
	RunE: func(cmd *cobra.Command, args []string) error {
		gitCmd := []string{"log", "--oneline"}
		if limit > 0 {
			gitCmd = append(gitCmd, "-n", strconv.Itoa(limit))
		}

		output, err := exec.Command("git", gitCmd...).Output()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				fmt.Printf("Error: %s\n", string(exitErr.Stderr))
			} else {
				fmt.Printf("Error executing git command: %v\n", err)
			}
			return nil
		}

		if len(output) == 0 {
			fmt.Println("No commit history found.")
			return nil
		}

		title := "Commit History"
		if limit > 0 {
			title = fmt.Sprintf("Latest %d Commits", limit)
		}

		fmt.Printf("\n%s\n", title)
		fmt.Println(string(output))

		return nil
	},
}

func init() {
	hisCmd.Flags().IntVarP(&limit, "limit", "n", 0, "Display only the latest n commit messages")
}
