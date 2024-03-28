package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Service struct {
	Name         string `yaml:"name"`
	Path         string `yaml:"path"`
	BuildCommand string `yaml:"build-command"`
	RunCommand   string `yaml:"run-command"`
	StopCommand  string `yaml:"stop-command"`
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
	replaceEnvVars(&config)
	return &config
}

func replaceEnvVars(config *Config) {
	for i, project := range config.Projects {
		// Update project path by env var, if present
		if isEnvVar(project.Path) {
			project.Path = os.Getenv(stripEnvChars(project.Path))
		}
		// Update each service's path by env var, if present
		for _, service := range project.Services {
			if isEnvVar(service.Path) {
				service.Path = os.Getenv(stripEnvChars(service.Path))
			}
		}
		// Persist changes in the config struct
		config.Projects[i] = project
	}

}

func isEnvVar(configValue string) bool {
	return strings.HasPrefix(configValue, "${") && strings.HasSuffix(configValue, "}")
}

func stripEnvChars(configValue string) string {
	return configValue[2 : len(configValue)-1]
}
