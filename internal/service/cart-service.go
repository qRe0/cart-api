package service

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/lib/pq"
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
	cart, err := c.repo.CreateCart()

	if err != nil {
		return nil, errs.ErrCreateCart
	}

	return cart, nil
}

func (c *CartService) AddItemToCart(cartIDStr string, item models.CartItem) (*models.CartItem, error) {
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil || cartID <= 0 {
		return nil, errs.ErrWrongCartID
	}
	if item.Product == "" {
		return nil, errs.ErrEmptyProductName
	}
	if item.Quantity <= 0 {
		return nil, errs.ErrWrongItemQuantity
	}

	item.CartID = cartID
	result, err := c.repo.AddItemToCart(item)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23503" {
				return nil, errs.ErrCartNotFound
			}

			return nil, fmt.Errorf("database error: %s", pqErr.Message)
		}
		return nil, err
	}

	return result, nil
}

func (c *CartService) RemoveItemFromCart(cartIDStr, itemIDStr string) error {
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil || cartID <= 0 {
		return errs.ErrWrongCartID
	}

	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil || itemID <= 0 {
		return errs.ErrWrongItemID
	}

	item := &models.CartItem{
		ID:     itemID,
		CartID: cartID,
	}

	err = c.repo.RemoveItemFromCart(item)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartService) GetCart(cartIDStr string) (*models.Cart, error) {
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil || cartID <= 0 {
		return nil, errs.ErrWrongCartID
	}

	cart := &models.Cart{
		ID: cartID,
	}

	resCart, err := c.repo.GetCart(cart)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.ErrCartNotFound
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	return resCart, nil
}
