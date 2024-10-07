package config

import (
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	DBPath          string
	Port            string
	LockKeys        []string
	DefaultTTL      time.Duration
	CleanupInterval time.Duration
}

// Load reads environment variables and returns a Config struct
func Load() *Config {
	lockKeysEnv := getEnv("LOCK_KEYS", "")
	var lockKeys []string
	if lockKeysEnv != "" {
		lockKeys = strings.Split(lockKeysEnv, ",")
		for i := range lockKeys {
			lockKeys[i] = strings.TrimSpace(lockKeys[i])
		}
	} else {
		log.Println("No LOCK_KEYS defined. All keys will be allowed.")
	}
	defaultTTL := valueAsDuration(
		getEnv("DEFAULT_TTL", ""),
		60*time.Second)
	cleanupInterval := valueAsDuration(
		getEnv("CLEANUP_INTERVAL", ""),
		30*time.Second)

	return &Config{
		DBPath:          getEnv("DB_PATH", "locks.db"),
		Port:            getEnv("PORT", "8080"),
		LockKeys:        lockKeys,
		DefaultTTL:      defaultTTL,
		CleanupInterval: cleanupInterval,
	}
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
