package mw

import (
	"net/http"

	"github.com/aiyaruch1320/go-login-jwt/login"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func OnlyAdmin(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*login.JwtCustomClaims)
		if claims.Role != "admin" {
			return c.JSON(http.StatusUnauthorized, bson.M{
				"message": "you are not authorized to access this resource",
			})
		}
		return h(c)
	}
}
