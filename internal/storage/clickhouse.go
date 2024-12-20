package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	
	"github.com/alphaedge-labs/tick-server/internal/config"
	"github.com/alphaedge-labs/tick-server/internal/models"
)

type ClickhouseStore struct {
	conn driver.Conn
}

func NewClickhouseStore(cfg *config.Config) (*ClickhouseStore, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", cfg.Clickhouse.Host, cfg.Clickhouse.Port)},
		Auth: clickhouse.Auth{
			Database: cfg.Clickhouse.Database,
			Username: cfg.Clickhouse.User,
			Password: cfg.Clickhouse.Password,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Debug: true,
	})
	
	if err != nil {
		return nil, fmt.Errorf("failed to connect to clickhouse: %v", err)
	}

	return &ClickhouseStore{conn: conn}, nil
}

func (s *ClickhouseStore) SaveTickData(ctx context.Context, data *models.TickData) error {
	query := `
		INSERT INTO tick_data (
			timestamp, symbol, expiry_date, product_type, right, strike_price,
			open, last, high, low, change, b_price, b_qty, s_price, s_qty,
			ltq, avg_price, oi, chng_oi, ttq, total_buy_qty, total_sell_qty,
			ttv, lower_ckt_lm, upper_ckt_lm, close, exchange, stock_name
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	return s.conn.Exec(ctx, query,
		data.Timestamp.Date,
		data.Data.Symbol,
		data.Metadata.ExpiryDate,
		data.Metadata.ProductType,
		data.Metadata.Right,
		data.Metadata.StrikePrice,
		data.Data.Open,
		data.Data.Last,
		data.Data.High,
		data.Data.Low,
		data.Data.Change,
		data.Data.BPrice,
		data.Data.BQty,
		data.Data.SPrice,
		data.Data.SQty,
		data.Data.Ltq,
		data.Data.AvgPrice,
		data.Data.OI,
		data.Data.CHNGOI,
		data.Data.Ttq,
		data.Data.TotalBuyQt,
		data.Data.TotalSellQ,
		data.Data.Ttv,
		data.Data.LowerCktLm,
		data.Data.UpperCktLm,
		data.Data.Close,
		data.Data.Exchange,
		data.Data.StockName,
	)
}

func (s *ClickhouseStore) SaveMarketDepth(ctx context.Context, depth *models.MarketDepth) error {
	query := `
		INSERT INTO market_depth (
			timestamp, symbol, expiry_date, product_type, right, strike_price,
			depth_level, buy_rate, buy_qty, buy_orders, sell_rate, sell_qty,
			sell_orders, stock_name
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	for i, level := range depth.Data.Depth {
		err := s.conn.Exec(ctx, query,
			depth.Timestamp.Date,
			depth.Data.Symbol,
			depth.Metadata.ExpiryDate,
			depth.Metadata.ProductType,
			depth.Metadata.Right,
			depth.Metadata.StrikePrice,
			i+1,
			level.BuyRate,
			level.BuyQty,
			level.BuyNoOfOrders,
			level.SellRate,
			level.SellQty,
			level.SellNoOfOrders,
			depth.Data.StockName,
		)
		if err != nil {
			return err
		}
	}
	return nil
} 