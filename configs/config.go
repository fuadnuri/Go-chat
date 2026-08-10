package config

import (
	"os"
	"time"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Host              string
	Port              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	ReadHeaderTimeout time.Duration
	IdleTimeout       time.Duration
	shutdownTimeout   time.Duration
}

func Load() Config {
	return Config{
		Server: ServerConfig{
			Host:              getEnv("SERVER_HOST", "0.0.0.0"),
			Port:              getEnv("SERVER_PORT", "8000"),
			ReadTimeout:       getDuration("SERVER_READ_TIMEOUT", 10*time.Second),
			WriteTimeout:      getDuration("SERVER_WRITE_TIMEOUT", 10*time.Second),
			ReadHeaderTimeout: getDuration("SERVER_READ_HEADER_TIMEOUT", 5*time.Second),
			IdleTimeout:       getDuration("SERVER_IDLE_TIMEOUT", 15*time.Second),
			shutdownTimeout:   getDuration("SERVER_SHUTDOWN_TIMEOUT", 10*time.Second),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getDuration(key string, fallback time.Duration) time.Duration {

	value := os.Getenv(key)

	if value != "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return duration
}
