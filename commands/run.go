package commands

import (
	"herder/config"
	"herder/utils"
	"log"
	"slices"
)

func StartProject(args *utils.ParsedArgs, config *config.Config) error {
	project, err := config.GetProject(args.Project)
	if err != nil {
		panic(err)
	}
	log.Printf("Starting services for project %s...\n", args.Project)
	for i, service := range project.Services {
		if slices.Contains(args.Exclude, service.Name) {
			log.Printf("\t%d. Service %s in -exclude list, skipping...", i+1, service.Name)
		} else if service.RunCommand == "" {
			log.Printf("\t%d. Service %s has no run command, skipping...", i+1, service.Name)
		} else {
			workDir := utils.GetFullPath(project.Path, service.Path)
			log.Printf("\t%d. Starting service %s...\n", i+1, service.Name)
			log.Printf("\t\tCommand: %s\n", service.RunCommand)
			log.Printf("\t\tDirectory: %s\n", workDir)
			out, err := utils.RunCommand(workDir, service.RunCommand)
			if err != nil {
				return err
			}
			log.Printf("\t\tOutput: %s\n", out)
		}
	}
	log.Println("Done")
	return nil
}
