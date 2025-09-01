//go:build !lambda
// +build !lambda

package main

import (
    "fmt"
    "log"

    "nexus-enedis-m23-measure-lambda/internal/config"
    "nexus-enedis-m23-measure-lambda/internal/router"
)

func main() {
    // Load .env files for local dev (no-op if absent)
    config.LoadDotEnv()
    cfg := config.Load()
    r := router.New()
    addr := fmt.Sprintf(":%d", cfg.Port)
    log.Printf("Starting local Gin server on http://localhost%s (env=%s)", addr, cfg.Env)
    if err := r.Run(addr); err != nil {
        log.Fatal(err)
    }
}
