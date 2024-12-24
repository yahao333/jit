package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/yahao333/jit/internal/ollama"
	"github.com/yahao333/jit/internal/utils"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available LLM models",
	RunE: func(cmd *cobra.Command, args []string) error {
		cm, err := utils.NewConfigManager()
		if err != nil {
			return err
		}

		if err := cm.EnsureConfig(); err != nil {
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

		// 创建表格
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Description", "Size", "Status", "Downloaded"})
		table.SetBorder(true)

		for modelName, details := range models {
			downloaded := "no"
			status := "disabled"

			if manager.IsModelValid(modelName) || details.Downloaded == "yes" {
				details.Downloaded = "yes"
				downloaded = "yes"
			}

			if modelName == config.ModelName {
				status = "active"
			}

			row := []string{
				modelName,
				details.Description,
				details.Size,
				status,
				downloaded,
			}
			table.Append(row)
		}

		if err := cm.SaveModels(models); err != nil {
			return err
		}

		table.Render()
		return nil
	},
}
