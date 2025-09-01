package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

// LoadDotEnv attempts to load environment variables from .env files for local dev.
// Order:
//  - .env (base defaults)
//  - .env.local (overrides for your machine)
// Missing files are ignored; OS env always wins over file values.
func LoadDotEnv() {
    // Load base .env if present
    if _, err := os.Stat(".env"); err == nil {
        if err := godotenv.Load(".env"); err != nil {
            log.Printf("[dotenv] failed loading .env: %v", err)
        }
    }
    // Load local overrides last
    if _, err := os.Stat(".env.local"); err == nil {
        if err := godotenv.Overload(".env.local"); err != nil {
            log.Printf("[dotenv] failed overloading .env.local: %v", err)
        }
    }
}

