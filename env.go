package configkit

import (
	"fmt"
	"os"
)

type osEnvRW struct {
}

func newOsEnvRW() EnvReaderWriter {
	return &osEnvRW{}
}

func (o *osEnvRW) Read(key string) string {
	return os.Getenv(key)
}

func (o *osEnvRW) ReadSafe(key string) (string, error) {
	val, isExist := os.LookupEnv(key)
	if !isExist {
		return "", fmt.Errorf("environment pair not found with key: %s", key)
	}
	return val, nil
}

func (o *osEnvRW) Write(key string, value string) error {
	return os.Setenv(key, value)
}
