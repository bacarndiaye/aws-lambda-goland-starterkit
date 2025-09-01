package config

import (
    "os"
    "strconv"
)

type Config struct {
    Env      string
    Port     int
    LogLevel string
}

func Load() Config {
    cfg := Config{
        Env:      getenv("ENV", "local"),
        Port:     getenvInt("PORT", 3000),
        LogLevel: getenv("LOG_LEVEL", "info"),
    }
    return cfg
}

func getenv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}

func getenvInt(key string, def int) int {
    if v := os.Getenv(key); v != "" {
        if n, err := strconv.Atoi(v); err == nil {
            return n
        }
    }
    return def
}

