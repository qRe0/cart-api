package service

import (
	"context"

	"github.com/qRe0/innowise-cart-api/internal/models"
)

type CartServiceInterface interface {
	CreateCart(ctx context.Context) (*models.Cart, error)
	AddItemToCart(cartIDStr string, item models.CartItem) (*models.CartItem, error)
	RemoveItemFromCart(cartIDStr, itemIDStr string) error
	GetCart(cartIDStr string) (*models.Cart, error)
}
