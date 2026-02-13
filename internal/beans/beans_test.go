package beans

import (
	"testing"
)

// This test requires the beans CLI to be installed and working in the environment.
// Since this is a live environment, we can actually try to use it.
func TestGetBean(t *testing.T) {
	client := NewClient()
	// do-vul should exist
	bean, err := client.GetBean("do-vul")
	if err != nil {
		t.Fatalf("Failed to get bean: %v", err)
	}

	if bean.ID != "do-vul" {
		t.Errorf("Expected bean ID do-vul, got %s", bean.ID)
	}
}

func TestIsBlocked(t *testing.T) {
	client := NewClient()
	blocked, err := client.IsBlocked("do-vul")
	if err != nil {
		t.Fatalf("Failed to check if blocked: %v", err)
	}

	// do-vul is likely not blocked by anything currently
	t.Logf("Bean do-vul blocked: %v", blocked)
}
