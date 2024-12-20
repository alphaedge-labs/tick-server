# Simple Tick Server

This is a simple tick server that simulates a real-time tick data stream. It's designed to be used for testing and development purposes.

## Prerequisites

-   Go 1.21 or higher
-   Docker and Docker Compose
-   ClickHouse database

## Setup and notes

1. First, make sure you have Go installed on your machine. You can check with:

```
go version
```

2. Clone the repository and navigate to the project directory:

```
git clone https://github.com/alphaedge-labs/tick-server.git
cd tick-server
```

3. Install Clickhouse on your local machine:

For Ubuntu/Debian:

```
# Add repository
sudo apt-get install -y apt-transport-https ca-certificates dirmngr
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 8919F6BD2B48D754

echo "deb https://packages.clickhouse.com/deb stable main" | sudo tee \
    /etc/apt/sources.list.d/clickhouse.list

# Install Clickhouse
sudo apt-get update
sudo apt-get install -y clickhouse-server clickhouse-client

# Start the service
sudo service clickhouse-server start
```

For MacOS:

```
# Using Homebrew
brew install clickhouse
brew services start clickhouse
```

4. Create a local .env file from the sample and modify for local development

```
cp .env.sample .env
```

5. Initialize the database schema:

```
# Connect to Clickhouse
clickhouse-client

# Run the migration script
cat migrations/init.sql | clickhouse-client
```

6. Install go dependencies

```
go mod download
```

7. Run the server:

```
go run cmd/server/main.go
```

To import data from a file, run the importer:

```
# Make the import script executable
chmod +x scripts/import_data.sh

# Run the import script
./scripts/import_data.sh
```

To test the WebSocket connection, you can use a tool like `websocat`:

```
# Install websocat
cargo install websocat

# Connect to the WebSocket server
websocat ws://localhost:8080/ws
```

Then send a subscription message:

```
{
    "action": "subscribe",
    "symbol": "4.1!41670",
    "throttle_ms": 100
}
```

To unsubscribe, send a message with the action set to "unsubscribe":

```
{
    "action": "unsubscribe",
    "symbol": "4.1!41670"
}
```

Useful commands for development:

1. Build the server:

```
go build -o server cmd/server/main.go
```

2. Run with debug logging:

```
DEBUG=1 go run cmd/server/main.go
```

3. Run with verbose logging:

```
VERBOSE=1 go run cmd/server/main.go
```
