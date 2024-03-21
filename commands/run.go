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
		log.Printf("\t%d. Starting service %s...\n", i+1, service.Name)
		log.Printf("\t\tCommand: %s\n", service.BuildCommand)
		log.Printf("\t\tDirectory: %s\n", service.Path)
		out, err := utils.RunCommand(service.Path, service.RunCommand)
		if err != nil {
			return err
		}
		log.Printf("\t\tOutput: %s\n", out)
	}
	log.Println("Done")
	return nil
}
