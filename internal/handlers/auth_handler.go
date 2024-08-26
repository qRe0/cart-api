package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	pb "github.com/qRe0/cart-api/proto/gen/go"
	"google.golang.org/grpc"
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

func (h *AuthHandler) LogIn(c *gin.Context) {
	// #TODO: Implement LogIn
}

func (h *AuthHandler) LogOut(c *gin.Context) {
	// #TODO: Implement LogOut
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	// #TODO: Implement Refresh
}
