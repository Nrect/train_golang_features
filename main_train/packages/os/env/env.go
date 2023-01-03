package env

import (
	"os"
)

func GetEnvValue(envName string) string {
	envValue, exists := os.LookupEnv(envName)
	if exists {
		return envValue
	}
	return ""
}

func SetEnvValue(envName, envValue string) error {
	err := os.Setenv(envName, envValue)
	if err != nil {
		return err
	}
	return nil
}
