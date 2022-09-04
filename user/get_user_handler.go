package user

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getUserFn func(context.Context, string) (*User, error)

func (gn getUserFn) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	return gn(ctx, username)
}

func GetUserHandler(svc getUserFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		user, err := svc.GetUserByUsername(c.Request().Context(), username)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}
