# jit

A CLI tool for generating Git commit messages using local LLM models powered by Ollama.

## Features

- 🤖 AI-powered commit message generation
- 🚀 Local LLM model support via Ollama
- 🔄 Multiple model management
- 📝 Interactive commit message editing
- 🌐 Cross-platform support

## Prerequisites

- Go 1.20 or higher
- Git
- [Ollama](https://ollama.ai/download)

## Installation 

```bash
go install github.com/yahao333/jit@latest
```

## Usage

1. Start the Ollama server:

```bash
jit start
```

2. Stage your changes:

```bash
git add .
```

3. Generate a commit message:

```bash
jit gen
```

4. (Optional) Generate and push:

```bash
jit gen -p
```

## Commands

- `start`: Start Ollama server and ensure default model is available
- `stop`: Stop the running Ollama server
- `gen`: Generate git commit message
- `list`: List all available LLM models
- `rm`: Delete a model from available models
- `use`: Select which model to use
- `his`: Display git commit history
- `version`: Print version information

## Configuration

Configuration files are stored in:
- Windows: `%APPDATA%/jit/`
- macOS: `~/Library/Application Support/jit/`
- Linux: `~/.config/jit/`

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

MIT License