package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yahao333/jit/internal/ollama"
	"github.com/yahao333/jit/internal/utils"
)

var useCmd = &cobra.Command{
	Use:   "use [model_name]",
	Short: "Select which model to use for generating commit messages",
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

		if _, exists := models[modelName]; !exists {
			fmt.Printf("Error: Unknown model '%s'\n", modelName)
			return listCmd.RunE(cmd, args)
		}

		pulled := true
		if models[modelName].Downloaded != "yes" {
			if err := manager.PullModel(modelName); err != nil {
				pulled = false
			}
		}

		if pulled {
			config, err := cm.GetConfig()
			if err != nil {
				return err
			}

			// 停用旧模型
			oldModel := models[config.ModelName]
			oldModel.Status = "disabled"
			models[config.ModelName] = oldModel

			// 激活新模型
			newModel := models[modelName]
			newModel.Status = "active"
			models[modelName] = newModel

			config.ModelName = modelName

			if err := cm.SaveConfig(config); err != nil {
				return err
			}
			if err := cm.SaveModels(models); err != nil {
				return err
			}

			fmt.Printf("Successfully switched to '%s' model.\n", modelName)
		} else {
			fmt.Println("\nDownload cancelled!")
		}

		return nil
	},
}
