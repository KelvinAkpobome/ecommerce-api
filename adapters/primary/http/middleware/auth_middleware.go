package middleware

import (
	"ecommerce-api/utils"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            c.JSON(http.StatusUnauthorized, utils.Response{
                Success: false,
                Message: "Missing or invalid Authorization header",
                Response: nil,
            })
            c.Abort()
            return
        }
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, utils.Response{
                Success: false,
                Message: "Invalid token",
                Response: nil,
            })
            c.Abort()
            return
        }

        c.Set("user_id", claims.UserID)
        c.Set("is_admin", claims.IsAdmin)
        c.Next()
    }
}

func IsAdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        isAdmin, exists := c.Get("is_admin")
        if !exists || !isAdmin.(bool) {
            c.JSON(http.StatusForbidden, utils.Response{
                Success: false,
                Message: "Access denied: Admins only",
                Response: nil,
            })
            c.Abort()
            return
        }
        c.Next()
    }
}
