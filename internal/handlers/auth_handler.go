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
	conn, err := grpc.Dial(address, grpc.WithBlock())
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

func (h *AuthHandler) SignUp(c *gin.Context) {
	// #TODO: Implement SignUp
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
