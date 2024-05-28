package service

import (
	"strconv"

	errs "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/models"
	"github.com/qRe0/innowise-cart-api/internal/repository"
)

type CartService struct {
	repo repository.CartRepositoryInterface
}

func NewCartService(repo repository.CartRepositoryInterface) *CartService {
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
		return nil, errs.ErrWrongCartID
	} else if cartID <= 0 {
		return nil, errs.ErrWrongCartID
	} else if item.Product == "" {
		return nil, errs.ErrEmptyProductName
	} else if item.Quantity <= 0 {
		return nil, errs.ErrWrongItemQuantity
	}

	return c.repo.AddItemToCart(cartID, item)
}

func (c *CartService) RemoveItemFromCart(cartIDStr, itemIDStr string) error {
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		return errs.ErrWrongCartID
	}
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		return errs.ErrWrongItemID
	}

	if cartID <= 0 {
		return errs.ErrWrongCartID
	} else if itemID <= 0 {
		return errs.ErrWrongItemID
	}

	return c.repo.RemoveItemFromCart(cartID, itemID)
}

func (c *CartService) GetCart(cartIDStr string) (*models.Cart, error) {
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		return nil, errs.ErrWrongCartID
	}
	if cartID <= 0 {
		return nil, errs.ErrWrongCartID
	}

	return c.repo.GetCart(cartID)
}
