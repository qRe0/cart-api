package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	errs "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/handlers"
	"github.com/qRe0/innowise-cart-api/internal/repository"
	"github.com/qRe0/innowise-cart-api/internal/service"
)

func Run() {
	db, err := repository.Init()
	if err != nil {
		log.Fatalln(err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			e := errs.ErrClosingDB
			log.Fatalln(e)
		}
	}(db)

	cartRepo := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepo)
	handler := handlers.NewHandler(cartService)

	http.HandleFunc("/carts", handler.CartHandler.CreateCart)
	http.HandleFunc("/carts/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.ItemHandler.AddItemToCart(w, r)
		case http.MethodGet:
			handler.CartHandler.GetCart(w, r)
		case http.MethodDelete:
			handler.ItemHandler.RemoveItemFromCart(w, r)
		}
	})

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load enviromental variables: %v", err)
	}

	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	log.Printf("Server is running on port %s", port)
	log.Fatalln(http.ListenAndServe(port, nil))
}
