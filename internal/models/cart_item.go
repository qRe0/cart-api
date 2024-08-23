package models

type CartItem struct {
	ID       int    `json:"id,omitempty"`
	CartID   int    `json:"cart_id,omitempty"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}
