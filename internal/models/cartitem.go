package models

type CartItem struct {
	Entity
	CartID   int    `json:"cart_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}
