package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/yahao333/jit/internal/ollama"
	"github.com/yahao333/jit/internal/utils"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start ollama server and ensure default model is available",
	RunE: func(cmd *cobra.Command, args []string) error {
		banner := `
╭──────────────────        Jit  ───────────────────╮
│            ⚡ AI-Powered Git Commits ⚡          │
│         Generated Locally, Commit Globally       │
╰──────────────────────────────────────────────────╯
`
		fmt.Println(banner)

		// 初始化配置
		cm, err := utils.NewConfigManager()
		if err != nil {
			return err
		}
		if err := cm.EnsureConfig(); err != nil {
			return err
		}

		// 检查ollama是否安装
		manager := ollama.NewOllamaManager("http://localhost:11434")
		if !manager.IsInstalled() {
			fmt.Println("Error: Ollama is not installed or not in PATH")
			fmt.Println("Please install Ollama following the instructions at: https://ollama.ai/download")
			return nil
		}

		// 检查服务是否运行
		if !manager.IsRunning() {
			if err := manager.StartService(); err != nil {
				return err
			}
			time.Sleep(3 * time.Second)
			fmt.Println("Ollama server started successfully!")
		} else {
			fmt.Println("Warning: Ollama server is already running!")
		}

		// 检查默认模型
		modelName := "llama3.2:3b"
		if !manager.IsModelValid(modelName) {
			fmt.Printf("Default model %s not found.\n", modelName)
			if err := manager.PullModel(modelName); err != nil {
				fmt.Println("Failed to pull default model. Please check your internet connection")
				return nil
			}
		} else {
			fmt.Printf("Default model %s is ready!\n", modelName)
		}

		return nil
	},
}
