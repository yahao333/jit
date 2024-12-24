package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type ModelConfig struct {
	Description string `json:"description"`
	Size        string `json:"size"`
	Status      string `json:"status"`
	Downloaded  string `json:"downloaded"`
}

type Config struct {
	ModelName string `json:"model_name"`
}

type ConfigManager struct {
	configDir  string
	cacheDir   string
	configFile string
	modelsFile string
}

// 默认模型配置
var DefaultModels = map[string]ModelConfig{
	"llama3.2:1b": {
		Description: "Lightweight model suitable for basic commit messages",
		Size:        "1.3GB",
		Status:      "disabled",
		Downloaded:  "no",
	},
	"gemma2:2b": {
		Description: "Enhanced lightweight model for moderate commit complexity",
		Size:        "1.6GB",
		Status:      "disabled",
		Downloaded:  "no",
	},
	"llama3.2:3b": {
		Description: "Balanced model for handling complex commits",
		Size:        "2.0GB",
		Status:      "disabled",
		Downloaded:  "no",
	},
	"llama3.1:8b": {
		Description: "High-performance model for detailed and intricate commits",
		Size:        "4.7GB",
		Status:      "disabled",
		Downloaded:  "no",
	},
}

// NewConfigManager 创建新的配置管理器
func NewConfigManager() (*ConfigManager, error) {
	configDir, err := getConfigDir()
	if err != nil {
		return nil, err
	}

	cacheDir, err := getCacheDir()
	if err != nil {
		return nil, err
	}

	cm := &ConfigManager{
		configDir:  configDir,
		cacheDir:   cacheDir,
		configFile: filepath.Join(configDir, "config.json"),
		modelsFile: filepath.Join(configDir, "models.json"),
	}

	return cm, nil
}

// getConfigDir 获取配置目录
func getConfigDir() (string, error) {
	var configBase string
	switch runtime.GOOS {
	case "windows":
		configBase = os.Getenv("APPDATA")
	case "darwin":
		configBase = filepath.Join(os.Getenv("HOME"), "Library", "Application Support")
	default: // linux and others
		configBase = filepath.Join(os.Getenv("HOME"), ".config")
	}

	configDir := filepath.Join(configBase, "jit")
	return configDir, nil
}

// getCacheDir 获取缓存目录
func getCacheDir() (string, error) {
	var cacheBase string
	switch runtime.GOOS {
	case "windows":
		cacheBase = os.Getenv("LOCALAPPDATA")
	case "darwin":
		cacheBase = filepath.Join(os.Getenv("HOME"), "Library", "Caches")
	default: // linux and others
		cacheBase = filepath.Join(os.Getenv("HOME"), ".cache")
	}

	cacheDir := filepath.Join(cacheBase, "jit")
	return cacheDir, nil
}

// EnsureConfig 确保配置文件存在
func (cm *ConfigManager) EnsureConfig() error {
	// 创建配置目录
	if err := os.MkdirAll(cm.configDir, 0755); err != nil {
		return err
	}

	// 初始化配置文件
	if _, err := os.Stat(cm.configFile); os.IsNotExist(err) {
		defaultConfig := Config{ModelName: "llama3.2:3b"}
		if err := cm.SaveConfig(defaultConfig); err != nil {
			return err
		}
	}

	// 初始化模型文件
	if _, err := os.Stat(cm.modelsFile); os.IsNotExist(err) {
		if err := cm.SaveModels(DefaultModels); err != nil {
			return err
		}
	}

	return nil
}

// GetConfig 获取当前配置
func (cm *ConfigManager) GetConfig() (Config, error) {
	if err := cm.EnsureConfig(); err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(cm.configFile)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

// GetModels 获取模型配置
func (cm *ConfigManager) GetModels() (map[string]ModelConfig, error) {
	if err := cm.EnsureConfig(); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(cm.modelsFile)
	if err != nil {
		return nil, err
	}

	var models map[string]ModelConfig
	if err := json.Unmarshal(data, &models); err != nil {
		return nil, err
	}

	return models, nil
}

// SaveConfig 保存配置
func (cm *ConfigManager) SaveConfig(config Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cm.configFile, data, 0644)
}

// SaveModels 保存模型配置
func (cm *ConfigManager) SaveModels(models map[string]ModelConfig) error {
	data, err := json.MarshalIndent(models, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(cm.modelsFile, data, 0644)
}

// GetActiveModel 获取当前活动模型
func (cm *ConfigManager) GetActiveModel() (string, error) {
	config, err := cm.GetConfig()
	if err != nil {
		return "", err
	}

	if config.ModelName == "" {
		return "llama3.2:3b", nil // 默认模型
	}

	return config.ModelName, nil
}

func (cm *ConfigManager) GetConfigFile() string {
	return cm.configFile
}
