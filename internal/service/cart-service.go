package service

import (
	"strconv"

	myErrors "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/models"
	r "github.com/qRe0/innowise-cart-api/internal/repo"
)

type CartService struct {
	repo r.ICartRepository
}

func NewCartService(repo r.ICartRepository) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (c *CartService) CreateCart() (*models.Cart, error) {
	return c.repo.CreateCart()
}

func (c *CartService) AddItemToCart(cartIDStr string, item models.CartItem) (*models.CartItem, error) {
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		return nil, myErrors.ErrWrongCartID
	} else if cartID <= 0 {
		return nil, myErrors.ErrWrongCartID
	} else if item.Product == "" {
		return nil, myErrors.ErrEmptyProductName
	} else if item.Quantity <= 0 {
		return nil, myErrors.ErrWrongItemQuantity
	}

	return c.repo.AddItemToCart(cartID, item)
}

func (c *CartService) RemoveItemFromCart(cartIDStr, itemIDStr string) error {
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		return myErrors.ErrWrongCartID
	}
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		return myErrors.ErrWrongItemID
	}

	if cartID <= 0 {
		return myErrors.ErrWrongCartID
	} else if itemID <= 0 {
		return myErrors.ErrWrongItemID
	}

	return c.repo.RemoveItemFromCart(cartID, itemID)
}

func (c *CartService) GetCart(cartIDStr string) (*models.Cart, error) {
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		return nil, myErrors.ErrWrongCartID
	}
	if cartID <= 0 {
		return nil, myErrors.ErrWrongCartID
	}

	return c.repo.GetCart(cartID)
}
