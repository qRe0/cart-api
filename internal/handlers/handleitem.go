package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"

	myErrors "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/models"
	s "github.com/qRe0/innowise-cart-api/internal/service"
)

type HandleItem struct {
	service *s.CartService
}

func NewHandleItem() *HandleItem {
	return &HandleItem{
		service: s.NewCartService(),
	}
}

func (h *HandleItem) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	cartIDStr := r.URL.Path[len("/carts/"):]
	cartIDStr = cartIDStr[:len(cartIDStr)-len("/items")]

	var parsedItem models.CartItem
	err := json.NewDecoder(r.Body).Decode(&parsedItem)
	if err != nil {
		http.Error(w, myErrors.ErrDecodingReqBody.Error(), http.StatusBadRequest)
		return
	}

	item, err := h.service.AddItemToCart(cartIDStr, parsedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, myErrors.ErrEncodingJSON.Error(), http.StatusBadRequest)
		return
	}
}

func (h *HandleItem) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	re := regexp.MustCompile(`/carts/(\d+)/items/(\d+)`)
	matches := re.FindStringSubmatch(path)

	var cartIDStr, itemIDStr string
	if len(matches) == 3 {
		cartIDStr = matches[1]
		itemIDStr = matches[2]
	} else {
		http.Error(w, myErrors.ErrInvalidURLFormat.Error(), http.StatusBadRequest)
	}

	err := h.service.RemoveItemFromCart(cartIDStr, itemIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("{}"))
	if err != nil {
		http.Error(w, myErrors.ErrEncodingJSON.Error(), http.StatusBadRequest)
		return
	}
}
