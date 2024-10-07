package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	DBPath          string
	Port            string
	LockKeys        []string
	DefaultTTL      time.Duration
	CleanupInterval time.Duration
}
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func valueAsDuration(valueString string, defaultValue time.Duration) time.Duration {
	if value, err := time.ParseDuration(valueString); err == nil {
		return value
	}
	log.Printf("Invalid duration %s. Using default %s.", valueString, defaultValue)
	return defaultValue
}
