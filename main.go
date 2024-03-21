package main

import (
	"herder/commands"
	"herder/config"
	"os"
)

func help() {
	// TODO: Print expected usage
}

func validateInput(project string, actions []string) error {
	// TODO: Verify the project exists in the config.
	// TODO: Verify all the actions are valid.
	return nil
}

func main() {
	args := os.Args
	project := args[1]
	actions := args[2:]
	// Read config
	configFilePath := "example.yml"
	config := config.ReadConfig(configFilePath)
	// Verify the project
	_, err := commands.VerifyServices("Demo Project", config)
	if err != nil {
		panic(err)
	}
	// Build the project
	err = commands.BuildProject("Demo Project", config)
	if err != nil {
		panic(err)
	}
	// Run the project
	err = commands.StartProject("Demo Project", config)
	if err != nil {
		panic(err)
	}

}
