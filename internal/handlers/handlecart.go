package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	myErrors "github.com/qRe0/innowise-cart-api/internal/errors"
	s "github.com/qRe0/innowise-cart-api/internal/service"
)

type HandleCart struct {
	service *s.CartService
}

func NewHandleCart() *HandleCart {
	return &HandleCart{
		service: s.NewCartService(),
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
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		http.Error(w, myErrors.ErrWrongCartID.Error(), http.StatusBadRequest)
		return
	}

	cart, err := h.service.GetCart(cartID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		http.Error(w, myErrors.ErrEncodingJSON.Error(), http.StatusBadRequest)
		return
	}
}
