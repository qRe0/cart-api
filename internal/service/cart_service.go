package service

import (
	"context"
	"strconv"

	errs "github.com/qRe0/cart-api/internal/errors"
	"github.com/qRe0/cart-api/internal/models"
	"github.com/qRe0/cart-api/internal/repository"
)

type CartService struct {
	repo repository.CartRepositoryInterface
}

func NewCartService(repo repository.CartRepositoryInterface) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (c *CartService) CreateCart(ctx context.Context) (*models.Cart, error) {
	cart, err := c.repo.CreateCart(ctx)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (c *CartService) AddItemToCart(ctx context.Context, cartIDStr string, item models.CartItem) (*models.CartItem, error) {
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
	result, err := c.repo.AddItemToCart(ctx, item)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CartService) RemoveItemFromCart(ctx context.Context, cartIDStr, itemIDStr string) error {
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

	err = c.repo.RemoveItemFromCart(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (c *CartService) GetCart(ctx context.Context, cartIDStr string) (*models.Cart, error) {
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil || cartID <= 0 {
		return nil, errs.ErrWrongCartID
	}

	cart := &models.Cart{
		ID: cartID,
	}

	resCart, err := c.repo.GetCart(ctx, cart)
	if err != nil {
		return nil, err
	}

	return resCart, nil
}
