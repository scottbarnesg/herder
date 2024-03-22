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
	Source       string `yaml:"source"`
}

type Project struct {
	Name     string    `yaml:"name"`
	Path     string    `yaml:"path"`
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

func (c *Config) HasProject(name string) bool {
	for _, project := range c.Projects {
		if project.Name == name {
			return true
		}
	}
	return false
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
