package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func GetPortBySerialID(serialID string) (string, error) {

	cmd := fmt.Sprintf("ls -l /dev/serial/by-id/ | grep '%s' | awk '{print $11}'", serialID)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	outStr := strings.TrimSpace(string(out))
	if outStr == "" {
		return "", errors.New("port not found")
	}

	pathArr := strings.Split(outStr, "/")
	port := pathArr[len(pathArr)-1]

	path := fmt.Sprintf("/dev/%s", port)

	return path, nil
}
