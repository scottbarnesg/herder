package config

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	configFilePath := "../example.yml"
	config := ReadConfig(configFilePath)
	log.Printf("%+v\n", config)
}

func TestReplaceEnvVars(t *testing.T) {
	// Set env vars
	os.Setenv("DEMO_PROJECT_PATH", "~/Documents/code/demo-project")
	// Read config
	configFilePath := "../example.yml"
	config := ReadConfig(configFilePath)
	// Verify env vars were read
	demoProject, _ := config.GetProject("DemoProject")
	assert.Equal(t, os.Getenv("DEMO_PROJECT_PATH"), demoProject.Path)
}
