package git

import (
	"os"
	"strings"
	"testing"
)

func TestGetRepoRoot(t *testing.T) {
	g := NewGit()
	root, err := g.GetRepoRoot()
	if err != nil {
		t.Fatalf("Failed to get repo root: %v", err)
	}

	wd, _ := os.Getwd()
	if !strings.Contains(wd, root) && !strings.Contains(root, wd) {
		t.Errorf("Repo root %s does not match working directory %s", root, wd)
	}
}

func TestCurrentBranch(t *testing.T) {
	g := NewGit()
	branch, err := g.CurrentBranch()
	if err != nil {
		t.Fatalf("Failed to get current branch: %v", err)
	}

	if branch == "" {
		t.Error("Current branch is empty")
	}
	t.Logf("Current branch: %s", branch)
}
