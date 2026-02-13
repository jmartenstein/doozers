package beans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Bean struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	BlockedBy []Bean `json:"blockedBy"`
}

type Client struct {
	Executable string
}

func NewClient() *Client {
	return &Client{
		Executable: "beans",
	}
}

func (c *Client) Query(query string) (map[string]interface{}, error) {
	cmd := exec.Command(c.Executable, "query", "--json", query)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("beans query failed: %w, stderr: %s", err, stderr.String())
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(out.Bytes(), &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal beans response: %w", err)
	}

	return resp, nil
}

func (c *Client) GetBean(id string) (*Bean, error) {
	query := fmt.Sprintf(`{ bean(id: "%s") { id title status blockedBy { id title status } } }`, id)
	resp, err := c.Query(query)
	if err != nil {
		return nil, err
	}

	beanData, ok := resp["bean"]
	if !ok || beanData == nil {
		return nil, fmt.Errorf("bean not found: %s", id)
	}

	beanJSON, err := json.Marshal(beanData)
	if err != nil {
		return nil, err
	}

	var bean Bean
	if err := json.Unmarshal(beanJSON, &bean); err != nil {
		return nil, err
	}

	return &bean, nil
}

func (c *Client) GetRootCause(id string) (*Bean, error) {
	bean, err := c.GetBean(id)
	if err != nil {
		return nil, err
	}

	if len(bean.BlockedBy) == 0 {
		return bean, nil
	}

	// For simplicity, we just follow the first blocker
	// In a real scenario, we might want to handle multiple blockers
	return c.GetRootCause(bean.BlockedBy[0].ID)
}

func (c *Client) IsBlocked(id string) (bool, error) {
	bean, err := c.GetBean(id)
	if err != nil {
		return false, err
	}

	return len(bean.BlockedBy) > 0, nil
}
