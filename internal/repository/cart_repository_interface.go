package repository

import (
	"context"

	"github.com/qRe0/innowise-cart-api/internal/models"
)

type CartRepositoryInterface interface {
	CreateCart(ctx context.Context) (*models.Cart, error)
	AddItemToCart(item models.CartItem) (*models.CartItem, error)
	RemoveItemFromCart(item *models.CartItem) error
	GetCart(cart *models.Cart) (*models.Cart, error)
}
