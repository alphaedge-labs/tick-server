CREATE DATABASE IF NOT EXISTS tickdata;

CREATE TABLE IF NOT EXISTS tickdata.tick_data
(
    timestamp DateTime64(3),
    symbol String,
    expiry_date String,
    product_type String,
    right String,
    strike_price String,
    open Float64,
    last Float64,
    high Float64,
    low Float64,
    change Float64,
    b_price Float64,
    b_qty Int64,
    s_price Float64,
    s_qty Int64,
    ltq Int64,
    avg_price Float64,
    oi String,
    chng_oi String,
    ttq Int64,
    total_buy_qty Int64,
    total_sell_qty Int64,
    ttv String,
    lower_ckt_lm Float64,
    upper_ckt_lm Float64,
    close Float64,
    exchange String,
    stock_name String
)
ENGINE = MergeTree()
ORDER BY (timestamp, symbol);

CREATE TABLE IF NOT EXISTS tickdata.market_depth
(
    timestamp DateTime64(3),
    symbol String,
    expiry_date String,
    product_type String,
    right String,
    strike_price String,
    depth_level UInt8,
    buy_rate Float64,
    buy_qty Int64,
    buy_orders Int64,
    sell_rate Float64,
    sell_qty Int64,
    sell_orders Int64,
    stock_name String
)
ENGINE = MergeTree()
ORDER BY (timestamp, symbol, depth_level); 