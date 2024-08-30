package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/qRe0/cart-api/docs"
	errs "github.com/qRe0/cart-api/internal/errors"
	"github.com/qRe0/cart-api/internal/models"
	"github.com/qRe0/cart-api/internal/service"
)

type ItemHandler struct {
	service service.CartServiceInterface
}

func NewItemHandler(cs service.CartServiceInterface) *ItemHandler {
	return &ItemHandler{
		service: cs,
	}
}

// AddItemToCart godoc
// @Tags CartItem Operations. This is a simple API for online shopping cart
// @Summary Add item to shopping cart
// @Schemes
// @Description This method allows user to add item to cart
// @Accept json
// @Produce json
// @Param cart_id path string true "Cart ID"
// @Param item body models.AddItemRequest true "Item"
// @Success 200 {object} models.CartItem "Item added to cart successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid cart ID"
// @Failure 401 {object} models.ErrorResponse "Empty claims"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /cart/{cart_id}/add [post]
// @Security ApiKeyAuth
func (h *ItemHandler) AddItemToCart(c *gin.Context) {
	ctx := c.Request.Context()

	cartIDStr := c.Param("cart_id")

	var parsedItem models.CartItem
	if err := c.ShouldBindJSON(&parsedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.ErrDecodingReqBody.Error()})
		return
	}

	item, err := h.service.AddItemToCart(ctx, cartIDStr, parsedItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// RemoveItemFromCart godoc
// @Tags CartItem Operations. This is a simple API for online shopping cart
// @Summary Remove item from shopping cart
// @Schemes
// @Description This method allows user to remove item from cart
// @Accept json
// @Produce json
// @Param cart_id path string true "Cart ID"
// @Param item_id path string true "Item ID"
// @Success 200 {object} models.RemoveItemMessageResponse "Item removed from cart successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid cart ID"
// @Failure 401 {object} models.ErrorResponse "Empty claims"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /cart/{cart_id}/remove/{item_id} [delete]
// @Security ApiKeyAuth
func (h *ItemHandler) RemoveItemFromCart(c *gin.Context) {
	ctx := c.Request.Context()

	cartIDStr := c.Param("cart_id")
	itemIDStr := c.Param("item_id")

	if cartIDStr == "" || itemIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL parameters"})
		return
	}

	if err := h.service.RemoveItemFromCart(ctx, cartIDStr, itemIDStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart successfully"})
}
