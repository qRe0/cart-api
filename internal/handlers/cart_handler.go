package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qRe0/innowise-cart-api/internal/service"
)

type CartHandler struct {
	service service.CartServiceInterface
}

func NewCartHandler(cs service.CartServiceInterface) *CartHandler {
	return &CartHandler{
		service: cs,
	}
}

func (h *CartHandler) CreateCart(c *gin.Context) {
	ctx := c.Request.Context()
	cart, err := h.service.CreateCart(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cart)
}

func (h *CartHandler) GetCart(c *gin.Context) {
	ctx := c.Request.Context()
	cartIDStr := c.Param("cart_id")

	cart, err := h.service.GetCart(ctx, cartIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}
