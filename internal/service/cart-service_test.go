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

func (m *MockCartRepository) RemoveItemFromCart(cartID, itemID int) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockCartRepository) IsCartExist(cartID int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockCartRepository) IsItemExist(itemID, cartID int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockCartRepository) GetCart(cartID int) (*models.Cart, error) {
	//TODO implement me
	panic("implement me")
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
			name:   "OK",
			cartID: "1",
			item: models.CartItem{
				Product:  "product",
				Quantity: 1,
			},
			setupMock: func() {
				mockRepo.On("AddItemToCart", mock.Anything).Return(
					&models.CartItem{
						Product:  "product",
						Quantity: 1,
						CartID:   1,
					},
					nil,
				)
			},
			expectedResult: &models.CartItem{
				Product:  "product",
				Quantity: 1,
				CartID:   1,
			},
			expectedErr: nil,
		},
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
