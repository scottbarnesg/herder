package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Service struct {
	Name         string `yaml:"name"`
	Path         string `yaml:"path"`
	BuildCommand string `yaml:"build-command"`
	RunCommand   string `yaml:"run-command"`
}

type Project struct {
	Name     string    `yaml:"name"`
	Services []Service `yaml:"services"`
}

type Config struct {
	Projects []Project `yaml:"projects"`
}

func (c *Config) GetProject(name string) (*Project, error) {
	for _, project := range c.Projects {
		if project.Name == name {
			return &project, nil
		}
	}
	return nil, fmt.Errorf("project name %s not in projects list", name)
}

func ReadConfig(filePath string) *Config {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}
	return &config
}
