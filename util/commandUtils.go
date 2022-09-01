package util

import (
	"os/exec"
)

func ExecuteCommand(command string, options ...string) (*string, error) {
	cmd := exec.Command(command, options...)

	if output, err := cmd.Output(); err != nil {
		return nil, err
	} else {
		result := string(output)
		return &result, nil
	}
}
