package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	pb "github.com/qRe0/cart-api/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthHandler struct {
	signUpClient  pb.SignUpClient
	logInClient   pb.LogInClient
	logOutClient  pb.LogOutClient
	refreshClient pb.RefreshClient
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

	return &AuthHandler{
		signUpClient:  signUpClient,
		logInClient:   logInClient,
		logOutClient:  logOutClient,
		refreshClient: refreshClient,
	}
}

// SignUp godoc
// @Tags Public routes. Registration and Authentication
// @Summary Sign-Up new user
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

	resp, err := h.signUpClient.SignUp(c, &user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

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

func (h *AuthHandler) Refresh(c *gin.Context) {
	// #TODO: Implement Refresh
}
