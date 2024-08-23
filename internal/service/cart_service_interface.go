package service

import (
	"context"

	"github.com/qRe0/cart-api/internal/models"
)

type CartServiceInterface interface {
	CreateCart(ctx context.Context) (*models.Cart, error)
	AddItemToCart(ctx context.Context, cartIDStr string, item models.CartItem) (*models.CartItem, error)
	RemoveItemFromCart(ctx context.Context, cartIDStr, itemIDStr string) error
	GetCart(ctx context.Context, cartIDStr string) (*models.Cart, error)
}
