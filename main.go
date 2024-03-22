package main

import (
	"fmt"
	"herder/commands"
	"herder/config"
	"os"
	"slices"
)

const helpCommand = "help"
const verifyCommand = "verify"
const cloneCommand = "clone"
const pullCommand = "pull"
const buildCommand = "build"
const runCommand = "run"
const stopCommand = "stop"

var acceptedActions = []string{helpCommand, verifyCommand, cloneCommand, pullCommand, buildCommand, runCommand, stopCommand}

func help() {
	// TODO: Print expected usage
}

func validateInput(projectName string, actions []string, config *config.Config) error {
	// Verify the project exists in the config.
	if !config.HasProject(projectName) {
		return fmt.Errorf("project %s is not defined", projectName)
	}
	// Verify all the actions are valid.
	for _, action := range actions {
		if !slices.Contains(acceptedActions, action) {
			return fmt.Errorf("invalid command provided: %s", action)
		}
	}
	return nil
}

func performAction(action string, projectName string, config *config.Config) error {
	if action == verifyCommand {
		err := commands.VerifyServices(projectName, config)
		if err != nil {
			return err
		}
	} else if action == buildCommand {
		err := commands.BuildProject(projectName, config)
		if err != nil {
			return err
		}
	} else if action == runCommand {
		err := commands.StartProject(projectName, config)
		if err != nil {
			return err
		}
	} else if action == cloneCommand {
		err := commands.CloneProjectRepos(projectName, config)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("action %s has not been implemented yet", action)
	}
	return nil
}

func main() {
	// Read input args
	args := os.Args
	projectName := args[1]
	actions := args[2:]
	// Read config
	configFilePath := "example.yml"
	config := config.ReadConfig(configFilePath)
	// Verify the inputs
	err := validateInput(projectName, actions, config)
	if err != nil {
		panic(err)
	}
	// Walk the provided actions and perform them
	for _, action := range actions {
		err := performAction(action, projectName, config)
		if err != nil {
			panic(err)
		}
	}

}
