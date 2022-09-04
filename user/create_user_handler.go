package user

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type createUserFn func(context.Context, *User) error

func (gn createUserFn) CreateUser(ctx context.Context, user *User) error {
	return gn(ctx, user)
}

func CreateUserHandler(svc getUserFn, svs createUserFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user *UserReq
		if err := c.Bind(&user); err != nil {
			return err
		}
		if _, err := svc.GetUserByUsername(c.Request().Context(), user.Username); err == nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "username already exists",
			})
		}
		if err := svs.CreateUser(c.Request().Context(), user.mapToUser()); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "user created",
		})
	}
}

func HashedPassword(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash)
}
