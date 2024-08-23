package models

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreationResponse struct {
	ID    string     `json:"id" example:"1"`
	Items []CartItem `json:"items"`
}

type AddItemRequest struct {
	Product  string `json:"product" example:"apple"`
	Quantity int    `json:"quantity" example:"10"`
}
