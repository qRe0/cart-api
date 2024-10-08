package repository

import (
	"context"

	"github.com/qRe0/innowise-cart-api/internal/models"
)

type CartRepositoryInterface interface {
	CreateCart(ctx context.Context) (*models.Cart, error)
	AddItemToCart(ctx context.Context, item models.CartItem) (*models.CartItem, error)
	RemoveItemFromCart(ctx context.Context, item *models.CartItem) error
	GetCart(ctx context.Context, cart *models.Cart) (*models.Cart, error)
}
