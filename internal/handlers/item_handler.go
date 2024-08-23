package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errs "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/models"
	"github.com/qRe0/innowise-cart-api/internal/service"
)

type ItemHandler struct {
	service service.CartServiceInterface
}

func NewItemHandler(cs service.CartServiceInterface) *ItemHandler {
	return &ItemHandler{
		service: cs,
	}
}

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
