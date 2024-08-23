package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/qRe0/innowise-cart-api/configs"
	errs "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/handlers"
	"github.com/qRe0/innowise-cart-api/internal/migrations"
	"github.com/qRe0/innowise-cart-api/internal/repository"
	"github.com/qRe0/innowise-cart-api/internal/service"
)

func Run() {
	cfg, err := configs.LoadEnv()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := repository.Init(cfg.DB)
	if err != nil {
		log.Fatalln(err)
	}

	migrator, err := migrations.NewMigrator(db)
	if err != nil {
		log.Fatalln(err)
	}
	err = migrator.Latest()
	if err != nil {
		log.Fatalln(err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(errs.ErrClosingDB)
		}
	}(db)

	cartRepository := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository)
	handler := handlers.NewHandler(cartService)

	router := gin.Default()

	cart := router.Group("/cart")
	cart.POST("/create", handler.CartHandler.CreateCart)
	cart.GET("/:cart_id/get", handler.CartHandler.GetCart)
	cart.POST("/:cart_id/add", handler.ItemHandler.AddItemToCart)
	cart.DELETE("/:cart_id/remove/:item_id:", handler.ItemHandler.RemoveItemFromCart)

	port := fmt.Sprintf(":%s", cfg.API.Port)
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("Server is running on port %s", port)
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	<-stop

	log.Println("Server is shutting down...")

	timeout, err := time.ParseDuration(cfg.API.ShutdownTimeout)
	if err != nil {
		log.Fatalf("Failed to parse shutdown timeout: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
