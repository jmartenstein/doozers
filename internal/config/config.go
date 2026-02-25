package config

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type SkillRef struct {
	Path string `yaml:"path"`
}

type SkillMetadata struct {
	Aliases        []string `yaml:"aliases"`
	UseGitWorktree bool     `yaml:"use_git_worktree"`
}

type SkillFrontmatter struct {
	Name        string        `yaml:"name"`
	Description string        `yaml:"description"`
	Metadata    SkillMetadata `yaml:"metadata"`
}

type Persona struct {
	Name           string   `yaml:"name"`
	Aliases        []string `yaml:"aliases"`
	Script         string   `yaml:"script"`
	Agent          string   `yaml:"agent"`
	UseGitWorktree bool     `yaml:"use_git_worktree"`
}

type Config struct {
	Version      string     `yaml:"version"`
	Skills       []SkillRef `yaml:"skills"`
	Personas     []Persona  `yaml:"personas"`
	DefaultAgent string     `yaml:"default_agent"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *Config) GetPersona(nameOrAlias string) (*Persona, bool) {
	// First check legacy personas
	for _, p := range c.Personas {
		if p.Name == nameOrAlias {
			return &p, true
		}
		for _, alias := range p.Aliases {
			if alias == nameOrAlias {
				return &p, true
			}
		}
	}

	// Then check new skills format
	for _, skillRef := range c.Skills {
		skillPath := filepath.Join(skillRef.Path, "SKILL.md")
		data, err := os.ReadFile(skillPath)
		if err != nil {
			continue
		}

		content := string(data)
		if !strings.HasPrefix(content, "---") {
			continue
		}

		parts := strings.SplitN(content, "---", 3)
		if len(parts) < 3 {
			continue
		}

		var fm SkillFrontmatter
		if err := yaml.Unmarshal([]byte(parts[1]), &fm); err != nil {
			continue
		}

		match := fm.Name == nameOrAlias
		if !match {
			for _, alias := range fm.Metadata.Aliases {
				if alias == nameOrAlias {
					match = true
					break
				}
			}
		}

		if match {
			return &Persona{
				Name:           fm.Name,
				Aliases:        fm.Metadata.Aliases,
				Script:         strings.TrimSpace(parts[2]),
				UseGitWorktree: fm.Metadata.UseGitWorktree,
			}, true
		}
	}

	return nil, false
}
