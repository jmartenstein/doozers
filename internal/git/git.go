package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Git struct{}

func NewGit() *Git {
	return &Git{}
}

func (g *Git) run(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("git %s failed: %w, stderr: %s", strings.Join(args, " "), err, stderr.String())
	}

	return strings.TrimSpace(out.String()), nil
}

func (g *Git) GetRepoRoot() (string, error) {
	return g.run("rev-parse", "--show-toplevel")
}

func (g *Git) CreateWorktree(path, branch string) error {
	_, err := g.run("worktree", "add", "-b", branch, path)
	return err
}

func (g *Git) RemoveWorktree(path string) error {
	_, err := g.run("worktree", "remove", path)
	return err
}

func (g *Git) CreateBranch(branch string) error {
	_, err := g.run("checkout", "-b", branch)
	return err
}

func (g *Git) Add(files ...string) error {
	args := append([]string{"add"}, files...)
	_, err := g.run(args...)
	return err
}

func (g *Git) Commit(message string) error {
	_, err := g.run("commit", "-m", message)
	return err
}

func (g *Git) Push(remote, branch string) error {
	_, err := g.run("push", remote, branch)
	return err
}

func (g *Git) CurrentBranch() (string, error) {
	return g.run("rev-parse", "--abbrev-ref", "HEAD")
}

func (g *Git) IsClean() (bool, error) {
	out, err := g.run("status", "--porcelain")
	if err != nil {
		return false, err
	}
	return out == "", nil
}
