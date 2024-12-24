package ollama

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/yahao333/jit/internal/errors"
)

type OllamaManager struct {
	baseURL string
}

func NewOllamaManager(baseURL string) *OllamaManager {
	return &OllamaManager{
		baseURL: baseURL,
	}
}

// IsInstalled 检查ollama是否已安装
func (m *OllamaManager) IsInstalled() bool {
	_, err := exec.LookPath("ollama")
	return err == nil
}

// IsRunning 检查ollama服务是否运行
func (m *OllamaManager) IsRunning() bool {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	_, err := client.Get(m.baseURL + "/api/tags")
	return err == nil
}

// StartService 启动ollama服务
func (m *OllamaManager) StartService() error {
	if !m.IsInstalled() {
		return errors.ErrOllamaNotInstalled
	}

	cmd := exec.Command("ollama", "serve")
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start ollama: %w", err)
	}

	// 等待服务启动
	for i := 0; i < 5; i++ {
		if m.IsRunning() {
			return nil
		}
		time.Sleep(time.Second)
	}

	return errors.ErrOllamaNotRunning
}

// StopService 停止ollama服务
func (m *OllamaManager) StopService() error {
	if !m.IsRunning() {
		return nil
	}

	cmd := exec.Command("pkill", "ollama")
	return cmd.Run()
}

// IsModelValid 检查模型是否存在
func (m *OllamaManager) IsModelValid(modelName string) bool {
	if !m.IsRunning() {
		return false
	}

	resp, err := http.Get(m.baseURL + "/api/tags")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var result struct {
		Models []struct {
			Name string `json:"name"`
		} `json:"models"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false
	}

	for _, model := range result.Models {
		if model.Name == modelName {
			return true
		}
	}
	return false
}

// PullModel 拉取模型
func (m *OllamaManager) PullModel(modelName string) error {
	if !m.IsRunning() {
		return errors.ErrOllamaNotRunning
	}

	cmd := exec.Command("ollama", "pull", modelName)
	if err := cmd.Run(); err != nil {
		return errors.ErrModelPullFailed
	}

	return nil
}

// DeleteModel 删除模型
func (m *OllamaManager) DeleteModel(modelName string) error {
	if !m.IsRunning() {
		return errors.ErrOllamaNotRunning
	}

	if !m.IsModelValid(modelName) {
		return errors.ErrModelNotFound
	}

	cmd := exec.Command("ollama", "rm", modelName)
	if err := cmd.Run(); err != nil {
		return errors.ErrModelDeleteFailed
	}

	return nil
}

/* 使用示例
// 创建manager实例
manager := ollama.NewOllamaManager("http://localhost:11434")

// 检查服务状态
if !manager.IsRunning() {
    // 启动服务
    if err := manager.StartService(); err != nil {
        log.Fatal(err)
    }
}

// 检查模型
if !manager.IsModelValid("codellama") {
    // 拉取模型
    if err := manager.PullModel("codellama"); err != nil {
        log.Fatal(err)
    }
}
*/
