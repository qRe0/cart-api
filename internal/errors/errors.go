package errors

import "errors"

var (
	ErrCreateCart        = errors.New("failed to create create cart")
	ErrWrongCartID       = errors.New("wrong cart id")
	ErrWrongItemID       = errors.New("wrong item id")
	ErrCartNotFound      = errors.New("cart not found")
	ErrItemNotFound      = errors.New("item not found")
	ErrEmptyProductName  = errors.New("empty product name")
	ErrWrongItemQuantity = errors.New("wrong item quantity")
	ErrClosingDB         = errors.New("failed to close database")
	ErrLoadEnvVars       = errors.New("failed to load environment variables")
	ErrDecodingReqBody   = errors.New("failed to decode request body")
	ErrStartTransaction  = errors.New("failed to start transaction")
	ErrCommitTransaction = errors.New("failed to commit transaction")
)
