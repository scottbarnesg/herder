package main

import (
	"fmt"
	"herder/commands"
	"herder/config"
	"herder/utils"
	"log"
	"os"
	"slices"
	"strings"
)

const helpCommand = "help"
const verifyCommand = "verify"
const cloneCommand = "clone"
const pullCommand = "pull"
const buildCommand = "build"
const runCommand = "run"
const stopCommand = "stop"

const includeFlag = "-include"
const excludeFlag = "-exclude"

var acceptedActions = []string{helpCommand, verifyCommand, cloneCommand, pullCommand, buildCommand, runCommand, stopCommand}
var acceptedKwargs = []string{includeFlag, excludeFlag}

func parseArgs(args []string) (*utils.ParsedArgs, error) {
	parsed := utils.ParsedArgs{Project: args[0]}
	i := 1
	for i < len(args) {
		arg := args[i]
		// If arg starts with "-", it is a kwarg.
		if strings.HasPrefix(arg, "-") {
			nextArg := args[i+1]
			i++
			if arg == includeFlag {
				parsed.Include = strings.Split(nextArg, ",")
			} else if arg == excludeFlag {
				parsed.Exclude = strings.Split(nextArg, ",")
			} else {
				return nil, fmt.Errorf("%s is an invalid flag. Valid flags are %+q", arg, acceptedKwargs)
			}
		} else { // Otherwise, it is a command
			parsed.Commands = append(parsed.Commands, arg)
		}
		i++
	}
	return &parsed, nil
}

func help() {
	// TODO: Print expected usage
}

func validateInput(args *utils.ParsedArgs, config *config.Config) error {
	// Verify the project exists in the config.
	if !config.HasProject(args.Project) {
		return fmt.Errorf("project %s is not defined", args.Project)
	}
	// Verify all the commands are valid.
	for _, action := range args.Commands {
		if !slices.Contains(acceptedActions, action) {
			return fmt.Errorf("invalid command provided: %s", action)
		}
	}
	return nil
}

func performActions(args *utils.ParsedArgs, config *config.Config) error {
	for _, action := range args.Commands {
		if action == verifyCommand {
			err := commands.VerifyServices(args, config)
			if err != nil {
				return err
			}
		} else if action == buildCommand {
			err := commands.BuildProject(args, config)
			if err != nil {
				return err
			}
		} else if action == runCommand {
			err := commands.StartProject(args, config)
			if err != nil {
				return err
			}
		} else if action == cloneCommand {
			err := commands.CloneProjectRepos(args, config)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("action %s has not been implemented yet", action)
		}
	}
	return nil
}

func main() {
	// Read input args
	args := os.Args[1:]
	parsed, err := parseArgs(args)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v\n", parsed)
	// Read config
	configFilePath := utils.ExpandPath("~/.herder/config.yml")
	config := config.ReadConfig(configFilePath)
	// Verify the inputs
	err = validateInput(parsed, config)
	if err != nil {
		panic(err)
	}
	// Run the commands
	err = performActions(parsed, config)
	if err != nil {
		panic(err)
	}
}
