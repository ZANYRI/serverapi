package env

import (
	"fmt"
	"os"

)

var ProviderToEnvKey = map[string]string{
    "GEMINI":   "GEMINI_KEY",
    "GIGACHAT": "GIGACHAT_KEY",
}

func GetAPIKey(provider string) (string, error) {
    envKey, ok := ProviderToEnvKey[provider]
    if !ok {
        return "", fmt.Errorf("unknown provider: %s", provider)
    }
    return GetEnvOrError(envKey)
}

func GetEnvOrError(v string) (string, error) {
    result := os.Getenv(v)
    if result == "" {
        return "", fmt.Errorf("environment variable %s is not set", v)
    }
    return result, nil
}
