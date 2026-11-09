package engine 

import "github.com/0zzy0sbourne/trading-analytics-engine/internal/domain"

type Orderbook struct {
	Symbol string 
	BuyOrders []*domain.Order 
	SellOrders []*domain.Order 
}

func NewOrderbook(symbol string) *Orderbook {
	orderbook := Orderbook{Symbol: symbol}
	return &orderbook
}

func Process(order *domain.Order) *domain.Trade {
	// gets a order, and creates a trade from it

	return nil // TODO: implement matching
}
