package handlers

import (
	"encoding/json"
	"net/http"

	errs "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/service"
)

type HandleCart struct {
	service service.ICartService
}

func NewHandleCart(cs service.ICartService) *HandleCart {
	return &HandleCart{
		service: cs,
	}
}

func (h *HandleCart) CreateCart(w http.ResponseWriter, r *http.Request) {
	cart, err := h.service.CreateCart()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *HandleCart) GetCart(w http.ResponseWriter, r *http.Request) {
	cartIDStr := r.URL.Path[len("/carts/"):]

	cart, err := h.service.GetCart(cartIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		http.Error(w, errs.ErrEncodingJSON.Error(), http.StatusBadRequest)
		return
	}
}
