FROM golang:1.21-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /tick-server ./cmd/server

# Create final lightweight image
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=0 /tick-server .
COPY .env .

EXPOSE 8080

CMD ["./tick-server"] 