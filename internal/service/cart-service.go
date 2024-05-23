package service

import (
	"github.com/qRe0/innowise-cart-api/internal/models"
	r "github.com/qRe0/innowise-cart-api/internal/repo"
)

type CartService struct {
	repo *r.CartRepository
}

func NewCartService() *CartService {
	return &CartService{
		repo: r.NewCartRepository(),
	}
}

func (c *CartService) CreateCart(cart models.Cart) error {
	return c.repo.CreateCart(cart)
}

func (c *CartService) AddItemToCart(cart models.Cart, item models.CartItem) error {
	return c.repo.AddItemToCart(cart, item)
}

func (c *CartService) RemoveItemFromCart(cart models.Cart, item models.CartItem) error {
	return c.repo.RemoveItemFromCart(cart, item)
}

func (c *CartService) GetCart(cart *models.Cart, item models.CartItem) error {
	return c.repo.GetCart(cart, item)
}

func (c *CartService) GetLastCartId() (int, error) {
	return c.repo.GetLastCartID()
}

func (c *CartService) GetLastItemID() (int, error) {
	return c.repo.GetLastItemID()
}
