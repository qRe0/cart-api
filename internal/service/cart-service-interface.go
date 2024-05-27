package service

import "github.com/qRe0/innowise-cart-api/internal/models"

type ICartService interface {
	CreateCart() (*models.Cart, error)
	AddItemToCart(cartIDStr string, item models.CartItem) (*models.CartItem, error)
	RemoveItemFromCart(cartIDStr, itemIDStr string) error
	GetCart(cartIDStr string) (*models.Cart, error)
}
