package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	pb "github.com/qRe0/cart-api/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthHandler struct {
	signUpClient     pb.SignUpClient
	logInClient      pb.LogInClient
	logOutClient     pb.LogOutClient
	refreshClient    pb.RefreshClient
	revokeClient     pb.RevokeClient
	MiddlewareClient pb.AuthMiddlewareClient
}

func NewAuthHandler(address string) *AuthHandler {
	log.Println("Connecting to gRPC server...")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	signUpClient := pb.NewSignUpClient(conn)
	logInClient := pb.NewLogInClient(conn)
	logOutClient := pb.NewLogOutClient(conn)
	refreshClient := pb.NewRefreshClient(conn)
	revokeClient := pb.NewRevokeClient(conn)
	middlewareClient := pb.NewAuthMiddlewareClient(conn)

	return &AuthHandler{
		signUpClient:     signUpClient,
		logInClient:      logInClient,
		logOutClient:     logOutClient,
		refreshClient:    refreshClient,
		revokeClient:     revokeClient,
		MiddlewareClient: middlewareClient,
	}
}

// SignUp godoc
// @Tags Public routes. Registration and Authentication
// @Summary Sign-Up new user and authorize him
// @Schemes
// @Description This method allows user to create a new user
// @Accept json
// @Produce json
// @Param user body models.SignUpRequest true "Input data for user registration"
// @Success 200 {object} models.SignUpResponse "User created successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input data"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /auth/signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var user pb.SignUpRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	md := metadata.MD{}
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	var metadataHeader metadata.MD
	resp, err := h.signUpClient.SignUp(ctx, &user, grpc.Header(&metadataHeader))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	accessToken := metadataHeader.Get("Authorization")
	refreshToken := metadataHeader.Get("Refresh-Token")
	if len(accessToken) == 0 || len(refreshToken) == 0 {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Header("Authorization", accessToken[0])
	c.Header("Refresh-Token", refreshToken[0])

	c.JSON(200, gin.H{"message": resp.Message})
}

// LogIn godoc
// @Tags Public routes. Registration and Authentication
// @Summary Log-In user
// @Schemes
// @Description This method allows user to log-in
// @Accept json
// @Produce json
// @Param user body models.LogInRequest true "Input data for user log-in"
// @Success 200 {object} models.LogInResponse "User logged-in successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input data"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /auth/login [post]
func (h *AuthHandler) LogIn(c *gin.Context) {
	var user pb.LogInRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	md := metadata.MD{}
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	var metadataHeader metadata.MD
	resp, err := h.logInClient.LogIn(ctx, &user, grpc.Header(&metadataHeader))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	accessToken := metadataHeader.Get("Authorization")
	refreshToken := metadataHeader.Get("Refresh-Token")
	if len(accessToken) == 0 || len(refreshToken) == 0 {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Header("Authorization", accessToken[0])
	c.Header("Refresh-Token", refreshToken[0])

	c.JSON(200, gin.H{"message": resp.Message})
}

func (h *AuthHandler) LogOut(c *gin.Context) {
	// #TODO: Implement LogOut
}

// Refresh godoc
// @Tags Public routes. Registration and Authentication
// @Summary Refresh user tokens (Access and Refresh)
// @Schemes
// @Description This method allows user to release new access and refresh tokens
// @Accept json
// @Produce json
// @Param token body models.RefreshRequest true "Input data for token refresh"
// @Success 200 {object} models.RefreshResponse "Tokens refreshed successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input data"
// @Failure 401 {object} models.ErrorResponse "Unauthorized"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var token pb.RefreshRequest
	err := c.ShouldBindJSON(&token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	md := metadata.MD{}
	ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

	var metadataHeader metadata.MD
	resp, err := h.refreshClient.Refresh(ctx, &token, grpc.Header(&metadataHeader))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	newAccessToken := metadataHeader.Get("Authorization")
	newRefreshToken := metadataHeader.Get("Refresh-Token")
	if len(newAccessToken) == 0 || len(newRefreshToken) == 0 {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Header("Authorization", newAccessToken[0])
	c.Header("Refresh-Token", newRefreshToken[0])

	c.JSON(200, gin.H{"message": resp.Message})
}

// RevokeTokens godoc
// @Tags Public routes. Registration and Authentication
// @Summary Revoke user tokens (Access and Refresh)
// @Schemes
// @Description This method allows user to revoke access and refresh tokens
// @Accept json
// @Produce json
// @Param token body models.RevokeRequest true "Input data for token revoke"
// @Success 200 {object} models.RevokeResponse "Tokens revoked successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid input data"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /auth/revoke [post]
func (h *AuthHandler) RevokeTokens(c *gin.Context) {
	var user pb.RevokeRequest
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.revokeClient.Revoke(c.Request.Context(), &user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": resp.Message})
}
