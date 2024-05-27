package errors

import "errors"

var (
	ErrConnectingToDB     = errors.New("error connecting to database")
	ErrGettingLastItemID  = errors.New("error getting last item id")
	ErrCreatingCart       = errors.New("error creating cart")
	ErrEncodingJSON       = errors.New("error encoding json")
	ErrWrongCartID        = errors.New("error. wrong cart id")
	ErrWrongItemID        = errors.New("error. wrong item id")
	ErrDecodingReqBody    = errors.New("error decoding request body")
	ErrAddItemToCart      = errors.New("error adding item to cart")
	ErrCartNotFound       = errors.New("error. cart not found in database")
	ErrGetItems           = errors.New("error. item not found in database")
	ErrRowsScan           = errors.New("error iterating rows with Scan() method")
	ErrItemNotFound       = errors.New("error. item not found in database")
	ErrRemoveItemFromCart = errors.New("error removing item from cart")
	ErrEmptyProductName   = errors.New("error. empty product name")
	ErrWrongItemQuantity  = errors.New("error. wrong item quantity")
	ErrInvalidURLFormat   = errors.New("error. invalid URL format")
	ErrClosingDB          = errors.New("error closing database")
)
