package utils

import (
	"os/exec"
	"strings"
)

func RunCommand(workDir string, command string) (string, error) {
	commandComponents := strings.Fields(command)
	cmd := exec.Command(commandComponents[0], commandComponents[1:]...)
	if workDir != "" {
		cmd.Dir = ExpandPath(workDir)
	}
	out, err := cmd.Output()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}
