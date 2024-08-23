package models

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateCartResponse struct {
	ID    int        `json:"id" example:"1"`
	Items []CartItem `json:"items"`
}

type GetCartResponse struct {
	ID    int        `json:"id" example:"1"`
	Items []CartItem `json:"items"`
}

type AddItemRequest struct {
	Product  string `json:"product" example:"apple"`
	Quantity int    `json:"quantity" example:"10"`
}

type AddItemResponse struct {
	ID       int    `json:"id" example:"1"`
	CartID   int    `json:"cart_id" example:"1"`
	Product  string `json:"product" example:"item1"`
	Quantity int    `json:"quantity" example:"1"`
}

type RemoveItemMessageResponse struct {
	Message string `json:"message" example:"Item removed from cart successfully"`
}
