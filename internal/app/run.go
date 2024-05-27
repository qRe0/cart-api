package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/qRe0/innowise-cart-api/internal/handlers"
)

func Run() {
	h := handlers.NewHandler()

	http.HandleFunc("/carts", h.HandleCart.CreateCart)
	http.HandleFunc("/carts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.HandleItem.AddItemToCart(w, r)
		case http.MethodGet:
			h.HandleCart.GetCart(w, r)
		case http.MethodDelete:
			h.HandleItem.RemoveItemFromCart(w, r)
		}
	})

	fmt.Println("Server is running on port :3000")
	log.Fatalln(http.ListenAndServe(":3000", nil))
}
