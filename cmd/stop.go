package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yahao333/jit/internal/ollama"
	"github.com/yahao333/jit/internal/utils"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running ollama server",
	RunE: func(cmd *cobra.Command, args []string) error {
		banner := `
╭──────────────────     Jit     ──────────────────╮
│          Local AI Models Are Resting 😴         │
│                  See You Soon! 🚀               │
╰─────────────────────────────────────────────────╯
`

		cm, err := utils.NewConfigManager()
		if err != nil {
			return err
		}

		models, err := cm.GetModels()
		if err != nil {
			return err
		}

		config, err := cm.GetConfig()
		if err != nil {
			return err
		}

		manager := ollama.NewOllamaManager("http://localhost:11434")
		if manager.IsRunning() {
			if err := manager.StopService(); err != nil {
				return err
			}

			// 更新模型状态
			activeModel := config.ModelName
			modelConfig := models[activeModel]
			modelConfig.Status = "disabled"
			models[activeModel] = modelConfig

			if err := cm.SaveConfig(config); err != nil {
				return err
			}
			if err := cm.SaveModels(models); err != nil {
				return err
			}

			// 删除配置文件
			os.Remove(cm.GetConfigFile())

			fmt.Println(banner)
			fmt.Println("Ollama server stopped successfully.")
		} else {
			fmt.Println("Warning: No Ollama server running found!")
		}

		return nil
	},
}
