package runner

import (
	"fmt"
	"os"
	"os/exec"
)

type Runner struct {
	Executable string
	Yolo       bool
}

func NewRunner(yolo bool) *Runner {
	return &Runner{
		Executable: "gemini",
		Yolo:       yolo,
	}
}

func (r *Runner) Run(prompt string, dir string) error {
	args := []string{"chat", "-p", prompt}
	if r.Yolo {
		args = append(args, "--yolo")
	}

	cmd := exec.Command(r.Executable, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("gemini execution failed: %w", err)
	}

	return nil
}
