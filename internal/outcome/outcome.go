package outcome

import (
	"fmt"
	"github.com/jmartenstein/doozers/internal/git"
)

type Manager struct {
	git *git.Git
}

func NewManager(g *git.Git) *Manager {
	return &Manager{
		git: g,
	}
}

func (m *Manager) Share(taskID, summary string) error {
	// For now, we add all changes
	if err := m.git.Add("."); err != nil {
		return err
	}

	clean, err := m.git.IsClean()
	if err != nil {
		return err
	}
	if clean {
		fmt.Println("No changes to commit")
		return nil
	}

	message := fmt.Sprintf("Task %s: %s", taskID, summary)
	if err := m.git.Commit(message); err != nil {
		return err
	}

	// Optionally push, but for now we might skip or make it configurable
	// branch, _ := m.git.CurrentBranch()
	// m.git.Push("origin", branch)

	return nil
}
