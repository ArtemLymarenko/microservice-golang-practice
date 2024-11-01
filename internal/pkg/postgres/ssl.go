package postgres

import commonconfig "project-management-system/internal/pkg/config"

const (
	sslDisable    = "disable"
	sslVerifyFull = "verify-full"
)

func getSSLConfig(env commonconfig.Env) string {
	var sslConfig string
	if env == commonconfig.EnvLocal {
		sslConfig = sslDisable
	} else {
		sslConfig = sslVerifyFull
	}

	return sslConfig
}
