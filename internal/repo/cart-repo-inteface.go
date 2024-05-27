package repository

import "github.com/qRe0/innowise-cart-api/internal/models"

type ICartRepository interface {
	CreateCart() (*models.Cart, error)
	AddItemToCart(cartID int, item models.CartItem) (*models.CartItem, error)
	RemoveItemFromCart(cartID, itemID int) error
	GetCart(cartID int) (*models.Cart, error)
}
