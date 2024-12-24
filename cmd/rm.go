package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yahao333/jit/internal/ollama"
	"github.com/yahao333/jit/internal/utils"
)

var rmCmd = &cobra.Command{
	Use:   "rm [model_name]",
	Short: "Delete a model from available models",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		modelName := args[0]

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

		models, err := cm.GetModels()
		if err != nil {
			return err
		}

		_, err = cm.GetConfig()
		if err != nil {
			return err
		}

		// 检查模型是否存在
		if !manager.IsModelValid(modelName) {
			fmt.Printf("Model %s doesn't exist, skipping deletion.\n", modelName)
			return nil
		}

		// 检查是否是当前活动模型
		if models[modelName].Status == "active" {
			fmt.Println("Error: Cannot remove currently active model!")
			fmt.Println("Please switch to a different model first using the 'use' command.")
			return nil
		}

		if models[modelName].Downloaded == "no" {
			fmt.Printf("Warning: Model: '%s' is not downloaded!\n", modelName)
		}

		return manager.DeleteModel(modelName)
	},
}
