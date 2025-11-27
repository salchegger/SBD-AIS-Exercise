package secrets

import (
	"errors"
	"fmt"
	"os"
)

const fileSuffix = "_FILE"

func LoadSecretOrEnv(envKey string) (string, error) {
	envVal, ok := os.LookupEnv(envKey)
	if ok {
		return envVal, nil
	}
	// lookup file envs
	envVal, ok = os.LookupEnv(envKey + fileSuffix)
	if !ok {
		return "", errors.New(fmt.Sprintf("environment variable '%s' or '%s_FILE'is not set", envKey, envKey))
	}
	// check if file exists
	_, err := os.Stat(envVal)
	if err != nil {
		return "", err
	}
	// load secret from file
	fileContent, err := os.ReadFile(envVal)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}
