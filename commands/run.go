package commands

import (
	"herder/config"
	"herder/utils"
	"log"
)

func StartProject(projectName string, config *config.Config) error {
	project, err := config.GetProject(projectName)
	if err != nil {
		panic(err)
	}
	log.Printf("Starting services for project %s...\n", projectName)
	for i, service := range project.Services {
		workDir := utils.GetFullPath(project.Path, service.Path)
		log.Printf("\t%d. Starting service %s...\n", i+1, service.Name)
		log.Printf("\t\tCommand: %s\n", service.BuildCommand)
		log.Printf("\t\tDirectory: %s\n", workDir)
		out, err := utils.RunCommand(workDir, service.RunCommand)
		if err != nil {
			return err
		}
		log.Printf("\t\tOutput: %s\n", out)
	}
	log.Println("Done")
	return nil
}
