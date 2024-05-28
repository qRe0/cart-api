package handlers

import s "github.com/qRe0/innowise-cart-api/internal/service"

type Handler struct {
	HandleCart *HandleCart
	HandleItem *HandleItem
}

func NewHandler(cs s.CartServiceInterface) *Handler {
	return &Handler{
		HandleCart: NewHandleCart(cs),
		HandleItem: NewHandleItem(cs),
	}
}
