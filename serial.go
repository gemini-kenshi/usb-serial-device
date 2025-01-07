package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func GetPortsBySerialID(serialID string) ([]string, error) {
	ports := []string{}

	cmd := fmt.Sprintf("ls -l /dev/serial/by-id/ | grep '%s' | awk '{print $11}'", serialID)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return ports, err
	}
	outStr := strings.TrimSpace(string(out))
	if outStr == "" {
		return ports, errors.New("port not found")
	}

	portStrs := strings.Split(outStr, "\n")
	for _, portStr := range portStrs {
		strs := strings.Split(portStr, "/")
		port := strs[len(strs)-1]
		ports = append(ports, fmt.Sprintf("/dev/%s", port))
	}

	return ports, nil
}

func GetCompleteSerialID(serialID string) ([]string, error) {
	cmd := fmt.Sprintf("ls -l /dev/serial/by-id/ | grep '%s' | awk '{print $9}'", serialID)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return []string{}, err
	}
	outStr := strings.TrimSpace(string(out))

	ids := strings.Split(outStr, "\n")
	return ids, nil
}
