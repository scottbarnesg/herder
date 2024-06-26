package utils

import (
	"os/user"
	"path/filepath"
	"strings"
)

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		usr, _ := user.Current()
		dir := usr.HomeDir
		path = filepath.Join(dir, path[2:])
	}
	return path
}

func GetFullPath(projectPath string, servicePath string) string {
	if projectPath != "" {
		return ExpandPath(filepath.Join(projectPath, servicePath))
	}
	return ExpandPath(servicePath)
}
