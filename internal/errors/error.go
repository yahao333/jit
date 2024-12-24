package errors

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

var (
	ErrNotGitRepo         = &CustomError{Code: 1001, Message: "not a git repository"}
	ErrNoChanges          = &CustomError{Code: 1002, Message: "no staged changes"}
	ErrOllamaAPI          = &CustomError{Code: 1003, Message: "ollama API error"}
	ErrOllamaNotInstalled = &CustomError{Code: 2001, Message: "ollama is not installed"}
	ErrOllamaNotRunning   = &CustomError{Code: 2002, Message: "ollama service is not running"}
	ErrModelNotFound      = &CustomError{Code: 2003, Message: "model not found"}
	ErrModelPullFailed    = &CustomError{Code: 2004, Message: "failed to pull model"}
	ErrModelDeleteFailed  = &CustomError{Code: 2005, Message: "failed to delete model"}
)
