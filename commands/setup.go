package commands

import (
	"herder/utils"
	"os"
)

func CreateProjectDir(projectDir string) error {
	expandedPath := utils.ExpandPath(projectDir)
	err := os.MkdirAll(expandedPath, os.ModePerm)
	return err
}
