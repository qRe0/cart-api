package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

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
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var item models.CartItem
	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	itemID, err := h.service.GetLastItemID()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusContinue)
	}

	item.CartID = cartID
	itemID++
	item.ID = &itemID

	cart := models.Cart{
		Entity: models.Entity{
			ID: &cartID,
		},
		Items: []models.CartItem{item},
	}

	err = h.service.AddItemToCart(cart, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *HandleItem) RemoveItemFromCart(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/carts/"):]
	cartIDStr := path[:len(path)-len("/items/")-1]
	itemIDStr := path[len(cartIDStr)+len("/items/"):]

	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var item models.CartItem
	item.ID = &itemID
	item.CartID = cartID

	cart := models.Cart{
		Entity: models.Entity{
			ID: &cartID,
		},
		Items: []models.CartItem{item},
	}

	err = h.service.RemoveItemFromCart(cart, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("{}"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
