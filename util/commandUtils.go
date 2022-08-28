package util

import (
	"os/exec"
)

func ExecuteCommand(command string, options ...string) (*string, error) {
	ls := exec.Command(command, options...)

	if output, err := ls.Output(); err != nil {
		return nil, err
	} else {
		result := string(output)
		return &result, nil
	}
}
