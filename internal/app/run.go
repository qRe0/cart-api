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
	"github.com/qRe0/cart-api/configs"
	errs "github.com/qRe0/cart-api/internal/errors"
	"github.com/qRe0/cart-api/internal/handlers"
	"github.com/qRe0/cart-api/internal/middleware"
	"github.com/qRe0/cart-api/internal/migrations"
	"github.com/qRe0/cart-api/internal/repository"
	"github.com/qRe0/cart-api/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	address := fmt.Sprintf("%s:%s", cfg.GRPC.Host, cfg.GRPC.Port)
	cartRepository := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository)
	handler := handlers.NewHandler(cartService, address)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	auth.POST("/signup", handler.AuthHandler.SignUp)
	auth.POST("/login", handler.AuthHandler.LogIn)
	auth.POST("/logout", handler.AuthHandler.LogOut)
	auth.POST("/refresh", handler.AuthHandler.Refresh)
	auth.POST("/revoke", handler.AuthHandler.RevokeTokens)

	cart := router.Group("/cart")
	cart.Use(middleware.AuthMiddleware(handler.AuthHandler.MiddlewareClient))
	{
		cart.POST("/create", handler.CartHandler.CreateCart)
		cart.GET("/:cart_id/get", handler.CartHandler.GetCart)
		cart.POST("/:cart_id/add", handler.ItemHandler.AddItemToCart)
		cart.DELETE("/:cart_id/remove/:item_id", handler.ItemHandler.RemoveItemFromCart)
	}

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
