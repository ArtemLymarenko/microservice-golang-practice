package appUtil

import "errors"

var ErrInvalidHttpPort = errors.New("invalid app port. must be 1025 < port < 51200")
