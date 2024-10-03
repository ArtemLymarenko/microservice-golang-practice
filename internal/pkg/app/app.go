package appUtil

import (
	"strconv"
	"strings"
)

func BuildHttpPath(host string, port int) (string, error) {
	minPort := 1025
	maxPort := 51200

	if port < minPort || port > maxPort {
		return "", ErrInvalidHttpPort
	}

	return strings.Join([]string{strings.Trim(host, " "), strconv.Itoa(port)}, ":"), nil
}
