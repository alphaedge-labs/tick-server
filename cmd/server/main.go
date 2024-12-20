package main

import (
    "log"

    "tick-server/internal/config"
    "tick-server/internal/server"
    "tick-server/internal/storage"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Initialize Clickhouse storage
    store, err := storage.NewClickhouseStore(cfg)
    if err != nil {
        log.Fatalf("Failed to initialize storage: %v", err)
    }

    // Create and start server
    srv := server.NewServer(cfg, store)
    if err := srv.Start(); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
} 