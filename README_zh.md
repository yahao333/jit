# jit

一个基于本地LLM模型的Git提交信息生成工具，由Ollama提供支持。

## 特性

- 🤖 AI驱动的提交信息生成
- 🚀 通过Ollama支持本地LLM模型
- 🔄 多模型管理
- 📝 交互式提交信息编辑
- 🌐 跨平台支持

## 前置要求

- Go 1.20 或更高版本
- Git
- [Ollama](https://ollama.ai/download)

## 安装 

```bash
go install github.com/yahao333/jit@latest
```

## 使用方法

1. 启动Ollama服务：

```bash
jit start
```

2. 暂存更改：

```bash
git add .
```

3. 生成提交信息：

```bash
jit gen
```

4. （可选）生成并推送：

```bash
jit gen -p
```

## 命令说明

- `start`: 启动Ollama服务并确保默认模型可用
- `stop`: 停止运行中的Ollama服务
- `gen`: 生成git提交信息
- `list`: 列出所有可用的LLM模型
- `rm`: 删除指定模型
- `use`: 选择要使用的模型
- `his`: 显示git提交历史
- `version`: 显示版本信息

## 配置

配置文件存储位置：
- Windows: `%APPDATA%/jit/`
- macOS: `~/Library/Application Support/jit/`
- Linux: `~/.config/jit/`

## 贡献指南

1. Fork 本仓库
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

MIT License