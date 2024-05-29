package errors

import "errors"

var (
	ErrCreatingCart      = errors.New("error creating cart")
	ErrWrongCartID       = errors.New("error with cart id (wrong cart id)")
	ErrWrongItemID       = errors.New("error with item id (wrong item id)")
	ErrCartNotFound      = errors.New("error finding cart id in database (cart not found)")
	ErrItemNotFound      = errors.New("error finding item id in database (item not found)")
	ErrEmptyProductName  = errors.New("error empty product name")
	ErrWrongItemQuantity = errors.New("error wrong item quantity")
	ErrClosingDB         = errors.New("error closing database")
)
