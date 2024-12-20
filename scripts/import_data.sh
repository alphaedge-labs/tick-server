#!/bin/bash

# Build the importer
go build -o importer ./cmd/importer

# Run the importer with the JSON files
./importer --ticks ./ticks-collection.json --depth ./market-depth-collection.json 