package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"

	errs "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/models"
	"github.com/qRe0/innowise-cart-api/internal/service"
)

type ItemHandler struct {
	service service.CartServiceInterface
}

func NewItemHandler(cs service.CartServiceInterface) *ItemHandler {
	return &ItemHandler{
		service: cs,
	}
}

func (h *ItemHandler) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cartIDStr := r.URL.Path[len("/carts/"):]
	cartIDStr = cartIDStr[:len(cartIDStr)-len("/items")]

	var parsedItem models.CartItem
	err := json.NewDecoder(r.Body).Decode(&parsedItem)
	if err != nil {
		http.Error(w, errs.ErrDecodingReqBody.Error(), http.StatusBadRequest)
		return
	}

	item, err := h.service.AddItemToCart(ctx, cartIDStr, parsedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, errs.ErrEncoding.Error(), http.StatusBadRequest)
		return
	}
}

func (h *ItemHandler) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	path := r.URL.Path
	re := regexp.MustCompile(`/carts/(\d+)/items/(\d+)`)
	matches := re.FindStringSubmatch(path)

	var cartIDStr, itemIDStr string
	if len(matches) == 3 {
		cartIDStr = matches[1]
		itemIDStr = matches[2]
	} else {
		http.Error(w, "error in URL", http.StatusBadRequest)
	}

	err := h.service.RemoveItemFromCart(ctx, cartIDStr, itemIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("{}"))
	if err != nil {
		http.Error(w, errs.ErrEncoding.Error(), http.StatusBadRequest)
		return
	}
}
