package handlers

import "github.com/qRe0/cart-api/internal/service"

type Handler struct {
	CartHandler *CartHandler
	ItemHandler *ItemHandler
	AuthHandler *AuthHandler
}

func NewHandler(serv service.CartServiceInterface, grpcAddress string) *Handler {
	return &Handler{
		CartHandler: NewCartHandler(serv),
		ItemHandler: NewItemHandler(serv),
		AuthHandler: NewAuthHandler(grpcAddress),
	}
}
