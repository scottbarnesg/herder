package commands

import (
	"herder/config"
	"herder/utils"
	"log"
	"slices"
)

func StopProject(args *utils.ParsedArgs, config *config.Config) error {
	project, err := config.GetProject(args.Project)
	if err != nil {
		panic(err)
	}
	log.Printf("Stopping services for project %s...\n", args.Project)
	for i, service := range project.Services {
		if slices.Contains(args.Exclude, service.Name) {
			log.Printf("\t%d. Service %s in -exclude list, skipping...", i+1, service.Name)
		} else if service.StopCommand == "" {
			log.Printf("\t%d. Service %s has no stop command, skipping...", i+1, service.Name)
		} else {
			workDir := utils.GetFullPath(project.Path, service.Path)
			log.Printf("\t%d. Stopping service %s...\n", i+1, service.Name)
			log.Printf("\t\tCommand: %s\n", service.StopCommand)
			log.Printf("\t\tDirectory: %s\n", workDir)
			out, err := utils.RunCommand(workDir, service.StopCommand)
			if err != nil {
				return err
			}
			log.Printf("\t\tOutput: %s\n", out)
		}
	}
	log.Println("Done")
	return nil
}
