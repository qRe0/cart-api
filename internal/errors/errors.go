package errors

import "errors"

var (
	ErrCreateCart        = errors.New("error can't create cart")
	ErrWrongCartID       = errors.New("error wrong cart id")
	ErrWrongItemID       = errors.New("error wrong item id")
	ErrCartNotFound      = errors.New("error cart not found")
	ErrItemNotFound      = errors.New("error item not found")
	ErrEmptyProductName  = errors.New("error empty product name")
	ErrWrongItemQuantity = errors.New("error wrong item quantity")
	ErrClosingDB         = errors.New("error closing database")
	ErrLoadEnvVars       = errors.New("error loading environment variables")
	ErrEncoding          = errors.New("error encoding")
	ErrDecodingReqBody   = errors.New("error decoding request body")
)
