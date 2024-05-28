package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	errs "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/handlers"
	"github.com/qRe0/innowise-cart-api/internal/repository"
	"github.com/qRe0/innowise-cart-api/internal/service"
)

func Run() {
	db, err := repository.Init()
	if err != nil {
		e := errs.ErrConnectingToDB
		fmt.Println(e)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			e := errs.ErrClosingDB
			panic(e)
		}
	}(db)

	cartRepo := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepo)
	handler := handlers.NewHandler(cartService)

	http.HandleFunc("/carts", handler.HandleCart.CreateCart)
	http.HandleFunc("/carts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.HandleItem.AddItemToCart(w, r)
		case http.MethodGet:
			handler.HandleCart.GetCart(w, r)
		case http.MethodDelete:
			handler.HandleItem.RemoveItemFromCart(w, r)
		}
	})

	fmt.Println("Server is running on port :3000")
	log.Fatalln(http.ListenAndServe(":3000", nil))
}
