package commands

import (
	"herder/config"
	"herder/utils"
	"log"
	"slices"
)

func BuildProject(args *utils.ParsedArgs, config *config.Config) error {
	project, err := config.GetProject(args.Project)
	if err != nil {
		return err
	}
	log.Printf("Building services for project %s...\n", args.Project)
	for i, service := range project.Services {
		if slices.Contains(args.Exclude, service.Name) {
			log.Printf("\t%d. Service %s in -exclude list, skipping...", i+1, service.Name)
		} else {
			workDir := utils.GetFullPath(project.Path, service.Path)
			log.Printf("\t%d. Building service %s...\n", i+1, service.Name)
			log.Printf("\t\tCommand: %s\n", service.BuildCommand)
			log.Printf("\t\tDirectory: %s\n", workDir)
			out, err := utils.RunCommand(workDir, service.BuildCommand)
			if err != nil {
				return err
			}
			log.Printf("\t\tOutput: %s\n", out)
		}
	}
	log.Println("Done")
	return nil
}
