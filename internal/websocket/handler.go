package websocket

import (
    "context"
    "encoding/json"
    "log"
    "sync"
    "time"

    "github.com/gorilla/websocket"
    
    "github.com/alphaedge-labs/tick-server/internal/config"
    "github.com/alphaedge-labs/tick-server/internal/models"
    "github.com/alphaedge-labs/tick-server/internal/storage"
)

type Subscription struct {
    Symbol     string
    Client     *websocket.Conn
    ThrottleMs int
}

type Handler struct {
    store         *storage.ClickhouseStore
    subscriptions sync.Map
    config        *config.Config
}

func NewHandler(store *storage.ClickhouseStore, cfg *config.Config) *Handler {
    return &Handler{
        store:  store,
        config: cfg,
    }
}

func (h *Handler) HandleWebSocket(conn *websocket.Conn) {
    defer conn.Close()

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Printf("Error reading message: %v", err)
            return
        }

        var msg struct {
            Action    string `json:"action"`
            Symbol    string `json:"symbol"`
            Throttle  int    `json:"throttle_ms,omitempty"`
        }

        if err := json.Unmarshal(p, &msg); err != nil {
            log.Printf("Error unmarshaling message: %v", err)
            continue
        }

        switch msg.Action {
        case "subscribe":
            throttle := msg.Throttle
            if throttle == 0 {
                throttle = h.config.Data.DefaultThrottleMs
            }
            
            sub := &Subscription{
                Symbol:     msg.Symbol,
                Client:     conn,
                ThrottleMs: throttle,
            }
            
            h.subscriptions.Store(conn.RemoteAddr().String()+msg.Symbol, sub)
            go h.streamData(sub)

        case "unsubscribe":
            h.subscriptions.Delete(conn.RemoteAddr().String() + msg.Symbol)

        default:
            log.Printf("Unknown action: %s", msg.Action)
        }
    }
}

func (h *Handler) streamData(sub *Subscription) {
    ticker := time.NewTicker(time.Duration(sub.ThrottleMs) * time.Millisecond)
    defer ticker.Stop()

    ctx := context.Background()

    for range ticker.C {
        // Check if subscription still exists
        if _, exists := h.subscriptions.Load(sub.Client.RemoteAddr().String() + sub.Symbol); !exists {
            return
        }

        // Stream both tick data and market depth
        if err := h.streamTickData(ctx, sub); err != nil {
            log.Printf("Error streaming tick data: %v", err)
        }

        if err := h.streamMarketDepth(ctx, sub); err != nil {
            log.Printf("Error streaming market depth: %v", err)
        }
    }
}

func (h *Handler) streamTickData(ctx context.Context, sub *Subscription) error {
    // Implementation for streaming tick data
    // This would typically query Clickhouse and send data to the websocket
    return nil
}

func (h *Handler) streamMarketDepth(ctx context.Context, sub *Subscription) error {
    // Implementation for streaming market depth
    // This would typically query Clickhouse and send data to the websocket
    return nil
} 