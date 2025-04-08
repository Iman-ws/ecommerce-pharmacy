package model

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Category   string  `json:"category"`
	StockLevel int     `json:"stock_level"`
	Price      float64 `json:"price"`
}
