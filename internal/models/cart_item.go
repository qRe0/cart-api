package models

type CartItem struct {
	ID       int    `json:"id" example:"1"`
	CartID   int    `json:"cart_id" example:"1"`
	Product  string `json:"product" example:"item1"`
	Quantity int    `json:"quantity" example:"1"`
}
