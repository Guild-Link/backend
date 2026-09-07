package common

import (
	"log"
	"os"
)

func MustEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalf("%s is not set", name)
	}
	return value
}

func DefaultEnv(name, defaultVal string) string {
	value := os.Getenv(name)
	if value == "" {
		return defaultVal
	}

	return value
}
