package login

import (
	"context"
	"net/http"

	"github.com/aiyaruch1320/go-login-jwt/user"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type getUserByUsername func(context.Context, string) (*user.User, error)

func (fn getUserByUsername) GetUserByUsername(ctx context.Context, username string) (*user.User, error) {
	return fn(ctx, username)
}

func LoginHandler(svc getUserByUsername) echo.HandlerFunc {
	return func(c echo.Context) error {
		var login Login
		if err := c.Bind(&login); err != nil {
			return err
		}
		user, err := svc.GetUserByUsername(c.Request().Context(), login.Username)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, bson.M{
				"message": "invalid username or password",
			})
		}
		if !checkPasswordHashed(login.Password, user.HashedPassword) {
			return c.JSON(http.StatusUnauthorized, bson.M{
				"message": "invalid username or password",
			})
		}
		// Set custom claims
		claims := &JwtCustomClaims{
			UserID:         user.ID.Hex(),
			Username:       user.Username,
			Role:           user.Role,
			StandardClaims: jwt.StandardClaims{},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, bson.M{
				"err": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, bson.M{
			"token": t,
		})
	}
}

func checkPasswordHashed(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
