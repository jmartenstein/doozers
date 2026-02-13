package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Persona struct {
	Name    string   `yaml:"name"`
	Aliases []string `yaml:"aliases"`
	Script  string   `yaml:"script"`
}

type Config struct {
	Personas []Persona `yaml:"personas"`
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
	return nil, false
}
