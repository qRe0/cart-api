package handlers

import "github.com/qRe0/innowise-cart-api/internal/service"

type Handler struct {
	CartHandler *CartHandler
	ItemHandler *ItemHandler
}

func NewHandler(cs service.CartServiceInterface) *Handler {
	return &Handler{
		CartHandler: NewCartHandler(cs),
		ItemHandler: NewItemHandler(cs),
	}
}
