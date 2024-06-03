package models

type CartItem struct {
	ID       int    `json:"id"`
	CartID   int    `json:"cart_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}
