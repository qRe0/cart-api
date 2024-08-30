package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	pb "github.com/qRe0/cart-api/proto/gen/go"
	"google.golang.org/grpc/metadata"
)

func AuthMiddleware(authClient pb.AuthMiddlewareClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}

		splitedAuthHeader := strings.Split(authHeader, " ")
		if len(splitedAuthHeader) != 2 || splitedAuthHeader[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		token := splitedAuthHeader[1]

		md := metadata.Pairs("Authorization", token)
		ctx := metadata.NewOutgoingContext(c.Request.Context(), md)

		resp, err := authClient.ValidateToken(ctx, &pb.ValidateTokenRequest{Token: token})
		//if err != nil || resp == nil || !resp.Valid {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		//	c.Abort()
		//	return
		//}
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if resp == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Response is nil"})
			c.Abort()
			return
		}
		if !resp.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		c.Set("uuid", resp.UserId)
		c.Next()
	}
}
