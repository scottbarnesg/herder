package commands

import (
	"fmt"
	"herder/config"
	"herder/utils"
	"os"
)

func VerifyServices(projectName string, config *config.Config) (bool, error) {
	project, err := config.GetProject(projectName)
	if err != nil {
		panic(err)
	}
	for _, service := range project.Services {
		if !exists(utils.ExpandPath(service.Path)) {
			return false, fmt.Errorf("path for service %s in project %s does not exist: %s", service.Name, projectName, service.Path)
		}
	}
	return true, nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
