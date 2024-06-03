package service

import (
	"testing"

	errs "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCartRepository struct {
	mock.Mock
}

func (m *MockCartRepository) CreateCart() (*models.Cart, error) {
	args := m.Called()
	return args.Get(0).(*models.Cart), args.Error(1)
}

func (m *MockCartRepository) AddItemToCart(item models.CartItem) (*models.CartItem, error) {
	args := m.Called(item)
	return args.Get(0).(*models.CartItem), args.Error(1)
}

func (m *MockCartRepository) RemoveItemFromCart(item *models.CartItem) error {
	args := m.Called(item)
	return args.Error(0)
}

func (m *MockCartRepository) GetCart(cart *models.Cart) (*models.Cart, error) {
	args := m.Called(cart)
	return args.Get(0).(*models.Cart), args.Error(1)
}

func TestCartService_CreateCart(t *testing.T) {
	mockRepo := new(MockCartRepository)
	service := NewCartService(mockRepo)

	expectedCart := &models.Cart{
		ID:    1,
		Items: []models.CartItem{},
	}
	mockRepo.On("CreateCart").Return(expectedCart, nil)

	cart, err := service.CreateCart()
	assert.NoError(t, err)
	assert.Equal(t, expectedCart, cart)
	mockRepo.AssertExpectations(t)
}

func TestCartService_AddItemToCart(t *testing.T) {
	mockRepo := new(MockCartRepository)
	service := NewCartService(mockRepo)

	testTable := []struct {
		name           string
		cartID         string
		item           models.CartItem
		setupMock      func()
		expectedResult *models.CartItem
		expectedErr    error
	}{
		{
			name:   "invalid cart ID",
			cartID: "0",
			item: models.CartItem{
				Product:  "product",
				Quantity: 1,
			},
			setupMock:      func() {},
			expectedResult: nil,
			expectedErr:    errs.ErrWrongCartID,
		},
		{
			name:   "empty product name",
			cartID: "1",
			item: models.CartItem{
				Product:  "",
				Quantity: 1,
			},
			setupMock:      func() {},
			expectedResult: nil,
			expectedErr:    errs.ErrEmptyProductName,
		},
		{
			name:   "wrong item quantity",
			cartID: "1",
			item: models.CartItem{
				Product:  "product",
				Quantity: 0,
			},
			setupMock:      func() {},
			expectedResult: nil,
			expectedErr:    errs.ErrWrongItemQuantity,
		},
		{
			name:   "cart not found",
			cartID: "100",
			item: models.CartItem{
				Product:  "product",
				Quantity: 1,
			},
			setupMock: func() {
				mockRepo.On("AddItemToCart", models.CartItem{CartID: 100, Product: "product", Quantity: 1}).Return(&models.CartItem{}, errs.ErrCartNotFound)
			},
			expectedResult: nil,
			expectedErr:    errs.ErrCartNotFound,
		},
		{
			name:   "OK",
			cartID: "1",
			item: models.CartItem{
				Product:  "product",
				Quantity: 1,
			},
			setupMock: func() {
				mockRepo.On("AddItemToCart", models.CartItem{CartID: 1, Product: "product", Quantity: 1}).Return(&models.CartItem{ID: 1, CartID: 1, Product: "product", Quantity: 1}, nil)
			},
			expectedResult: &models.CartItem{ID: 1, CartID: 1, Product: "product", Quantity: 1},
			expectedErr:    nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupMock()

			result, err := service.AddItemToCart(testCase.cartID, testCase.item)
			if testCase.expectedErr != nil {
				assert.ErrorIs(t, err, testCase.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, testCase.expectedResult, result)

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestCartService_RemoveItemFromCart(t *testing.T) {
	mockRepo := new(MockCartRepository)
	service := NewCartService(mockRepo)

	testTable := []struct {
		name        string
		cartID      string
		itemID      string
		setupMock   func()
		expectedErr error
	}{
		{
			name:        "invalid cart ID",
			cartID:      "0",
			itemID:      "1",
			setupMock:   func() {},
			expectedErr: errs.ErrWrongCartID,
		},
		{
			name:        "invalid item ID",
			cartID:      "1",
			itemID:      "0",
			setupMock:   func() {},
			expectedErr: errs.ErrWrongItemID,
		},
		{
			name:   "cart not found",
			cartID: "100",
			itemID: "1",
			setupMock: func() {
				mockRepo.On("RemoveItemFromCart", &models.CartItem{ID: 1, CartID: 100}).Return(errs.ErrCartNotFound)
			},
			expectedErr: errs.ErrCartNotFound,
		},
		{
			name:   "item not found",
			cartID: "1",
			itemID: "100",
			setupMock: func() {
				mockRepo.On("RemoveItemFromCart", &models.CartItem{ID: 100, CartID: 1}).Return(errs.ErrItemNotFound)
			},
			expectedErr: errs.ErrItemNotFound,
		},
		{
			name:   "OK",
			cartID: "1",
			itemID: "1",
			setupMock: func() {
				mockRepo.On("RemoveItemFromCart", &models.CartItem{ID: 1, CartID: 1}).Return(nil)
			},
			expectedErr: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupMock()

			err := service.RemoveItemFromCart(testCase.cartID, testCase.itemID)
			if testCase.expectedErr != nil {
				assert.ErrorIs(t, err, testCase.expectedErr)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestCartService_GetCart(t *testing.T) {
	mockRepo := new(MockCartRepository)
	service := NewCartService(mockRepo)

	testTable := []struct {
		name           string
		cartID         string
		setupMock      func()
		expectedResult *models.Cart
		expectedErr    error
	}{
		{
			name:           "invalid cart ID",
			cartID:         "0",
			setupMock:      func() {},
			expectedResult: nil,
			expectedErr:    errs.ErrWrongCartID,
		},
		{
			name:   "cart not found",
			cartID: "100",
			setupMock: func() {
				mockRepo.On("GetCart", &models.Cart{ID: 100}).Return(&models.Cart{}, errs.ErrCartNotFound)
			},
			expectedResult: nil,
			expectedErr:    errs.ErrCartNotFound,
		},
		{
			name:   "OK",
			cartID: "1",
			setupMock: func() {
				mockRepo.On("GetCart", mock.Anything).Return(
					&models.Cart{
						ID: 1,
						Items: []models.CartItem{
							{
								ID:       1,
								CartID:   1,
								Product:  "product",
								Quantity: 1,
							},
						},
					},
					nil,
				)
			},
			expectedResult: &models.Cart{
				ID: 1,
				Items: []models.CartItem{
					{
						ID:       1,
						CartID:   1,
						Product:  "product",
						Quantity: 1,
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupMock()

			result, err := service.GetCart(testCase.cartID)
			if testCase.expectedErr != nil {
				assert.ErrorIs(t, err, testCase.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, testCase.expectedResult, result)

			mockRepo.AssertExpectations(t)
		})
	}
}
