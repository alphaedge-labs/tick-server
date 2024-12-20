package models

import "time"

type Metadata struct {
    ExpiryDate  string `json:"expiry_date"`
    ProductType string `json:"product_type"`
    Right       string `json:"right"`
    StrikePrice string `json:"strike_price"`
    Symbol      string `json:"symbol"`
}

type TickData struct {
    Timestamp struct {
        Date time.Time `json:"$date"`
    } `json:"timestamp"`
    Metadata Metadata `json:"metadata"`
    Data     struct {
        Symbol      string  `json:"symbol"`
        Open        float64 `json:"open"`
        Last        float64 `json:"last"`
        High        float64 `json:"high"`
        Low         float64 `json:"low"`
        Change      float64 `json:"change"`
        BPrice      float64 `json:"bPrice"`
        BQty        int64   `json:"bQty"`
        SPrice      float64 `json:"sPrice"`
        SQty        int64   `json:"sQty"`
        Ltq         int64   `json:"ltq"`
        AvgPrice    float64 `json:"avgPrice"`
        Quotes      string  `json:"quotes"`
        OI          string  `json:"OI"`
        CHNGOI      string  `json:"CHNGOI"`
        Ttq         int64   `json:"ttq"`
        TotalBuyQt  int64   `json:"totalBuyQt"`
        TotalSellQ  int64   `json:"totalSellQ"`
        Ttv         string  `json:"ttv"`
        LowerCktLm  float64 `json:"lowerCktLm"`
        UpperCktLm  float64 `json:"upperCktLm"`
        Close       float64 `json:"close"`
        Exchange    string  `json:"exchange"`
        StockName   string  `json:"stock_name"`
    } `json:"data"`
    ID string `json:"_id"`
} 