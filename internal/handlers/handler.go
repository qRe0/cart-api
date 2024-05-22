package handlers

type Handler struct {
	HandleCart *HandleCart
	HandleItem *HandleItem
}

func NewHandler() *Handler {
	return &Handler{
		HandleCart: NewHandleCart(),
		HandleItem: NewHandleItem(),
	}
}
