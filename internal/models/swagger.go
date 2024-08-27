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

type SignUpRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Phone    string `json:"phone" binding:"required" example:"+1111111111"`
	Email    string `json:"email" binding:"required" example:"johndoe1@gmail.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type SignUpResponse struct {
	Message string `json:"message" example:"User created successfully!"`
}

type LogInRequest struct {
	Phone    string `json:"phone" binding:"required" example:"+1111111111"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LogInResponse struct {
	Message string `json:"message" example:"User logged-in successfully!"`
}
