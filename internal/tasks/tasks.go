package tasks

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/rscardinho/golang-dotfiles/internal/config"
)

type Status struct {
	Elapsed   time.Duration
	StdoutStr string
	StderrStr string
	ExitCode  int
}

var (
	green  = "\033[32m"
	red    = "\033[31m"
	purple = "\033[35m"
	none   = "\033[0m"
)

func ExecuteAll(title string, tasks []config.Task, logFile *os.File) {
	fmt.Printf("%s\n\n", title)

	for _, task := range tasks {
		Execute(task.Name, task.Script, task.Validation, logFile)
	}
}

func Execute(name, script string, validationScript string, logFile *os.File) {
	fmt.Printf(" %s…%s  %s", purple, none, name)

	start := time.Now()

	logEntry := fmt.Sprintf("\n--- [%s] %s ---\n", name, start.Format("2006-01-02 15:04:05"))
	_, _ = logFile.WriteString(logEntry)

	fullScript := script
	if validationScript != "" {
		fullScript += "\n" + validationScript
	}
	fullScript = "set -e\n" + fullScript

	status, err := CheckStatus(start, fullScript)

	logEntry = fmt.Sprintf(
		"Elapsed: %s\nScript:\n%s\n\nStdout:\n%s\n\nStderr:\n%s\n\nError: %v (Exit code: %d)\n",
		status.Elapsed, fullScript, status.StdoutStr, status.StderrStr, err, status.ExitCode,
	)
	_, _ = logFile.WriteString(logEntry)

	if err != nil {
		fmt.Printf("\r %s✘%s  %s\n", red, none, name)
	} else {
		fmt.Printf("\r %s✔%s  %s\n", green, none, name)
	}
}

func CheckStatus(start time.Time, script string) (Status, error) {
	cmd := exec.Command("bash", "-c", script)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()

	elapsed := time.Since(start)

	stdoutStr := stdoutBuf.String()
	stderrStr := stderrBuf.String()

	var exitCode int
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			exitCode = -1
		}
	}

	status := Status{
		Elapsed:   elapsed,
		StdoutStr: stdoutStr,
		StderrStr: stderrStr,
		ExitCode:  exitCode,
	}

	return status, err
}
