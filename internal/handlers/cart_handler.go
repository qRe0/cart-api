package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/qRe0/cart-api/docs"
	"github.com/qRe0/cart-api/internal/service"
)

type CartHandler struct {
	service service.CartServiceInterface
}

func NewCartHandler(cs service.CartServiceInterface) *CartHandler {
	return &CartHandler{
		service: cs,
	}
}

// CreateCart godoc
// @Tags Cart Operations. This is a simple API for online shopping cart
// @Summary Creates shopping cart
// @Schemes
// @Description This method allows user to create cart
// @Accept json
// @Produce json
// @Success 201 {object} models.Cart
// @Failure 401 {object} models.ErrorResponse "Empty claims"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /cart/create [post]
// @Security ApiKeyAuth
func (h *CartHandler) CreateCart(c *gin.Context) {
	ctx := c.Request.Context()
	cart, err := h.service.CreateCart(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cart)
}

// GetCart godoc
// @Tags Cart Operations. This is a simple API for online shopping cart
// @Summary Get shopping cart
// @Schemes
// @Description This method allows user to get cart by ID
// @Accept json
// @Produce json
// @Param cart_id path string true "Cart ID"
// @Success 200 {object} models.Cart
// @Failure 400 {object} models.ErrorResponse "Invalid cart ID"
// @Failure 401 {object} models.ErrorResponse "Empty claims"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /cart/{cart_id}/get [get]
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
