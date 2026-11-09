package domain

import (
	"time"
)

// order side 
type OrderSide string 
// order type 
type OrderType string 

const (
	// define order side enums
	Buy OrderSide = "BUY" 
	Sell OrderSide = "SELL" 

	// define order type enums 
	Limit OrderType = "LIMIT" 
	Market OrderType = "MARKET"
)

type Order struct {
	ID string 
	UserID string
	Symbol string
	Side OrderSide
	Type OrderType
	Price float64
	Quantity float64 
	Timestamp time.Time
}

type Trade struct { 
	TradeID string 
	BuyOrderID string 
	SellOrderID string 
	Symbol string
	Price float64
	Quantity float64
	ExecutedAt time.Time 
}



