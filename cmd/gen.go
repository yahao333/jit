package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yahao333/jit/internal/git"
	"github.com/yahao333/jit/internal/ollama"
	"github.com/yahao333/jit/internal/utils"
)

var pushFlag bool

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate git commit message",
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := ollama.NewOllamaManager("http://localhost:11434")
		if !manager.IsRunning() {
			fmt.Println("Error: Ollama server is not running!")
			fmt.Println("Start the ollama server by running 'jit start' command.")
			return nil
		}

		cm, err := utils.NewConfigManager()
		if err != nil {
			return err
		}

		config, err := cm.GetConfig()
		if err != nil {
			return err
		}

		client := ollama.NewClient("http://localhost:11434", config.ModelName)
		commitManager := git.NewCommitManager(client)

		if err := commitManager.Generate(); err != nil {
			return err
		}

		if pushFlag {
			fmt.Println("Pushing changes to the remote repository...")
			if err := git.Push(); err != nil {
				fmt.Printf("Error: Auto pushing FAILED! %v\n", err)
				return nil
			}
			fmt.Println("Push successful!")
		}

		return nil
	},
}

func init() {
	genCmd.Flags().BoolVarP(&pushFlag, "push", "p", false, "Enable auto-push")
}
