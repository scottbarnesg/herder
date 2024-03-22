package commands

import (
	"fmt"
	"herder/config"
	"herder/utils"
	"log"
)

func CloneProjectRepos(projectName string, config *config.Config) error {
	// Get the project
	project, err := config.GetProject(projectName)
	if err != nil {
		return err
	}
	// Create the project path if it doesn't exist
	if project.Path != "" {
		err := CreateProjectDir(project.Path)
		if err != nil {
			return err
		}
	}
	// For each service, clone the repo to the target directory
	log.Printf("Cloning repos for project %s...\n", projectName)
	for _, service := range project.Services {
		targetDir := utils.GetFullPath(project.Path, service.Path)
		if PathExists(targetDir) {
			log.Printf("Directory %s already exists, skipping clone of %s...\n", targetDir, service.Source)
			continue
		}
		log.Printf("Cloning %s to %s\n", service.Source, targetDir)
		_, err := cloneRepository(service.Source, targetDir)
		if err != nil {
			return err
		}
	}
	log.Println("Done.")
	return nil
}

func cloneRepository(repoUrl string, targetDir string) (string, error) {
	// Expand targetDir
	expandedPath := utils.ExpandPath(targetDir)
	// Execute command
	commandString := fmt.Sprintf("git clone %s %s", repoUrl, expandedPath)
	return utils.RunCommand("", commandString)
}

func PullReposForProject() {
	// TODO: Implement
}

func PullRepository() {
	// TODO: Implement
}
