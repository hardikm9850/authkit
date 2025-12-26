package middleware

import (
	jwt "github.com/hardikm9850/authkit/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "userID"
	ContextRolesKey  = "roles"
)

func JWTAuth(jwtManager jwt.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtManager.VerifyAccessToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set(ContextUserIDKey, claims.UserID)
		c.Set(ContextRolesKey, claims.Roles)
		c.Next()
	}
}
