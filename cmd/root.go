package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yahao333/jit/internal/version"
)

var rootCmd = &cobra.Command{
	Use:     "jit",
	Short:   "A CLI tool for git commit message generation",
	Version: version.Version(),
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(genCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(useCmd)
	rootCmd.AddCommand(hisCmd)
}
