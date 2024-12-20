package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/alphaedge-labs/tick-server/internal/config"
	"github.com/alphaedge-labs/tick-server/internal/models"
	"github.com/alphaedge-labs/tick-server/internal/storage"
)

func main() {
	tickDataFile := flag.String("ticks", "", "Path to ticks-collection.json file")
	marketDepthFile := flag.String("depth", "", "Path to market-depth-collection.json file")
	flag.Parse()

	if *tickDataFile == "" || *marketDepthFile == "" {
		log.Fatal("Both --ticks and --depth file paths are required")
	}

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

	// Import tick data
	if err := importTickData(*tickDataFile, store); err != nil {
		log.Fatalf("Failed to import tick data: %v", err)
	}

	// Import market depth data
	if err := importMarketDepth(*marketDepthFile, store); err != nil {
		log.Fatalf("Failed to import market depth data: %v", err)
	}
}

func importTickData(filePath string, store *storage.ClickhouseStore) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open tick data file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	ctx := context.Background()
	count := 0

	// Read opening bracket
	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("failed to read opening bracket: %v", err)
	}

	// Read array elements
	for decoder.More() {
		var tickData models.TickData
		if err := decoder.Decode(&tickData); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to decode tick data: %v", err)
		}

		if err := store.SaveTickData(ctx, &tickData); err != nil {
			return fmt.Errorf("failed to save tick data: %v", err)
		}

		count++
		if count%1000 == 0 {
			log.Printf("Imported %d tick data records", count)
		}
	}

	log.Printf("Successfully imported %d tick data records", count)
	return nil
}

func importMarketDepth(filePath string, store *storage.ClickhouseStore) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open market depth file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	ctx := context.Background()
	count := 0

	// Read opening bracket
	_, err = decoder.Token()
	if err != nil {
		return fmt.Errorf("failed to read opening bracket: %v", err)
	}

	// Read array elements
	for decoder.More() {
		var marketDepth models.MarketDepth
		if err := decoder.Decode(&marketDepth); err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to decode market depth: %v", err)
		}

		if err := store.SaveMarketDepth(ctx, &marketDepth); err != nil {
			return fmt.Errorf("failed to save market depth: %v", err)
		}

		count++
		if count%1000 == 0 {
			log.Printf("Imported %d market depth records", count)
		}
	}

	log.Printf("Successfully imported %d market depth records", count)
	return nil
} 