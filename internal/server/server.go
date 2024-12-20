package server

import (
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
    
    "github.com/alphaedge-labs/tick-server/internal/config"
    "github.com/alphaedge-labs/tick-server/internal/storage"
    "github.com/alphaedge-labs/tick-server/internal/websocket"
)

type Server struct {
    cfg     *config.Config
    store   *storage.ClickhouseStore
    wsHandler *websocket.Handler
    upgrader websocket.Upgrader
}

func NewServer(cfg *config.Config, store *storage.ClickhouseStore) *Server {
    upgrader := websocket.Upgrader{
        ReadBufferSize:  cfg.Websocket.ReadBufferSize,
        WriteBufferSize: cfg.Websocket.WriteBufferSize,
        CheckOrigin: func(r *http.Request) bool {
            return true // In production, implement proper origin checking
        },
    }

    return &Server{
        cfg:      cfg,
        store:    store,
        wsHandler: websocket.NewHandler(store, cfg),
        upgrader:  upgrader,
    }
}

func (s *Server) Start() error {
    http.HandleFunc("/ws", s.handleWebSocket)
    
    addr := fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)
    fmt.Printf("Server starting on %s\n", addr)
    
    return http.ListenAndServe(addr, nil)
}

func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := s.upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Printf("Failed to upgrade connection: %v\n", err)
        return
    }

    s.wsHandler.HandleWebSocket(conn)
} 