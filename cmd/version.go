package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yahao333/jit/internal/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("jit version %s\n", version.Version())
	},
}
