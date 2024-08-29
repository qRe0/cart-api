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

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	errs "github.com/qRe0/cart-api/internal/errors"
	"github.com/qRe0/cart-api/internal/handlers"
	"github.com/qRe0/cart-api/internal/migrations"
	"github.com/qRe0/cart-api/internal/models"
	"github.com/qRe0/cart-api/internal/repository"
	"github.com/qRe0/cart-api/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	var cfg models.Config
	err := env.Parse(&cfg)

	db, err := repository.Init(cfg)
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	cart := router.Group("/cart")
	cart.POST("/create", handler.CartHandler.CreateCart)
	cart.GET("/:cart_id/get", handler.CartHandler.GetCart)
	cart.POST("/:cart_id/add", handler.ItemHandler.AddItemToCart)
	cart.DELETE("/:cart_id/remove/:item_id", handler.ItemHandler.RemoveItemFromCart)

	port := fmt.Sprintf(":%s", cfg.APIPort)
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

	timeout := cfg.ShutdownTimeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}
