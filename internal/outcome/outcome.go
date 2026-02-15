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

	branch, err := m.git.CurrentBranch()
	if err != nil {
		return fmt.Errorf("failed to get current branch: %w", err)
	}

	fmt.Printf("Pushing changes to origin %s...\n", branch)
	if err := m.git.Push("origin", branch); err != nil {
		return fmt.Errorf("failed to push changes: %w", err)
	}

	return nil
}
