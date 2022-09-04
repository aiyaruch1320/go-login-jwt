package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type createUserFn func(context.Context, *User) error

func (gn createUserFn) CreateUser(ctx context.Context, user *User) error {
	return gn(ctx, user)
}

func (u *UserReq) Validate() error {
	if u.Role != RoleAdmin && u.Role != RoleUser {
		return errors.New("role must be admin or user")
	}
	return nil
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
		if err := user.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
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
