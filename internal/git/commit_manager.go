package git

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"

	"github.com/yahao333/jit/internal/errors"
	"github.com/yahao333/jit/internal/ollama"
)

type CommitManager struct {
	ollamaClient *ollama.Client
}

func NewCommitManager(ollamaClient *ollama.Client) *CommitManager {
	return &CommitManager{
		ollamaClient: ollamaClient,
	}
}

// executeGitCommand 执行git命令
func (m *CommitManager) executeGitCommand(args []string) (string, error) {
	cmd := exec.Command("git", args...)

	// Windows特定处理
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow: true,
		}
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("git command failed: %w", err)
	}

	return string(output), nil
}

// CheckStagedChanges 检查staged的更改
func (m *CommitManager) CheckStagedChanges() (string, error) {
	output, err := m.executeGitCommand([]string{"diff", "--staged"})
	if err != nil {
		return "", err
	}

	if strings.TrimSpace(output) == "" {
		return "", errors.ErrNoChanges
	}

	return output, nil
}

// GenerateCommitMessage 生成commit消息
func (m *CommitManager) GenerateCommitMessage(diff string) (string, error) {
	systemPrompt := `You are a Git expert specializing in concise and meaningful commit messages based on output of git diff command.
    Choose a type from below that best describes the git diff output:
        fix: A bug fix,
        docs: Documentation only changes,
        style: Changes that do not affect the meaning of the code,
        refactor: A code change that neither fixes a bug nor adds a feature,
        perf: A code change that improves performance,
        test: Adding missing tests or correcting existing tests,
        build: Changes that affect the build system or external dependencies,
        ci: Changes to our CI configuration files and scripts,
        chore: Other changes that don't modify src or test files,
        revert: Reverts a previous commit,
        feat: A new feature.
    Generate a concise git commit message written in present tense in the format "type: description".
    Maximum length 40 characters, no explanations.`

	prompt := fmt.Sprintf("%s\n\nGit diff:\n%s", systemPrompt, diff)
	message, err := m.ollamaClient.Generate(prompt)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(message), nil
}

// EditCommitMessage 编辑commit消息
func (m *CommitManager) EditCommitMessage(initialMessage string) (string, error) {
	fmt.Printf("Initial commit message: %s\n", initialMessage)
	fmt.Print("Edit message (press Enter to keep, or type new message): ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return initialMessage, nil
	}
	return input, nil
}

// PerformGitCommit 执行git commit
func (m *CommitManager) PerformGitCommit(message string) error {
	_, err := m.executeGitCommand([]string{"commit", "-m", message})
	return err
}

// Generate 执行完整的commit生成流程
func (m *CommitManager) Generate() error {
	// 检查staged changes
	diff, err := m.CheckStagedChanges()
	if err != nil {
		return err
	}

	// 生成commit消息
	initialMessage, err := m.GenerateCommitMessage(diff)
	if err != nil {
		return err
	}

	// 编辑commit消息
	finalMessage, err := m.EditCommitMessage(initialMessage)
	if err != nil {
		return err
	}

	// 执行commit
	return m.PerformGitCommit(finalMessage)
}

/*
使用示例
client := ollama.NewClient("http://localhost:11434", "codellama")
manager := git.NewCommitManager(client)
err := manager.Generate()
if err != nil {
    log.Fatal(err)
}
*/
