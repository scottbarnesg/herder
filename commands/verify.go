package commands

import (
	"fmt"
	"herder/config"
	"herder/utils"
	"log"
	"os"
)

func VerifyServices(projectName string, config *config.Config) error {
	project, err := config.GetProject(projectName)
	if err != nil {
		return err
	}
	for _, service := range project.Services {
		log.Printf("Verifying service %s...\n", service.Name)
		if !PathExists(utils.GetFullPath(project.Path, service.Path)) {
			return fmt.Errorf("path for service %s in project %s does not exist: %s", service.Name, projectName, service.Path)
		}
	}
	log.Println("OK")
	return nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
