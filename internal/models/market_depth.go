package models

import "time"

type DepthLevel struct {
    BuyRate       float64 `json:"BestBuyRate"`
    BuyQty        int64   `json:"BestBuyQty"`
    BuyNoOfOrders int64   `json:"BuyNoOfOrders"`
    BuyFlag       string  `json:"BuyFlag"`
    SellRate      float64 `json:"BestSellRate"`
    SellQty       int64   `json:"BestSellQty"`
    SellNoOfOrders int64  `json:"SellNoOfOrders"`
    SellFlag      string  `json:"SellFlag"`
}

type MarketDepth struct {
    Timestamp struct {
        Date time.Time `json:"$date"`
    } `json:"timestamp"`
    Metadata Metadata `json:"metadata"`
    Data     struct {
        Symbol      string       `json:"symbol"`
        Time        string       `json:"time"`
        Depth       []DepthLevel `json:"depth"`
        Quotes      string       `json:"quotes"`
        StockName   string       `json:"stock_name"`
    } `json:"data"`
    ID string `json:"_id"`
} 