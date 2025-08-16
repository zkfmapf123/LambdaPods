package internal

import (
	"os"
	"strings"
)

func GetEnvReturnDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

func MatchStringEquals(str string, targets ...string) bool {
	for _, t := range targets {
		if str == t {
			return true
		}
	}

	return false
}

func MatchStringContains(str string, targets ...string) bool {
	for _, t := range targets {
		if strings.Contains(str, t) {
			return true
		}
	}

	return false
}
