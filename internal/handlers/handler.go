package handlers

import "github.com/qRe0/innowise-cart-api/internal/service"

type Handler struct {
	HandleCart *HandleCart
	HandleItem *HandleItem
}

func NewHandler(cs service.CartServiceInterface) *Handler {
	return &Handler{
		HandleCart: NewHandleCart(cs),
		HandleItem: NewHandleItem(cs),
	}
}
