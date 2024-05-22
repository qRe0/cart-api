package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/qRe0/innowise-cart-api/internal/models"
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
	var cart models.Cart
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	cartId, err := h.service.GetLastCartId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cartId++
	cart.ID = &cartId
	cart.Items = []models.CartItem{}

	err = h.service.CreateCart(cart)
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

}
