package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aiyaruch1320/go-login-jwt/config"
	"github.com/aiyaruch1320/go-login-jwt/login"
	"github.com/aiyaruch1320/go-login-jwt/mw"
	"github.com/aiyaruch1320/go-login-jwt/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := config.InitMongoDB(ctx)
	defer func() {
		if err := db.Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mongodb := db.Client.Database("test")

	e := echo.New()
	// config
	corsConfig := middleware.CORSConfig{
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}
	e.Use(middleware.CORSWithConfig(corsConfig))

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.POST("/login", login.LoginHandler(user.GetUserByUsername(mongodb)))
	e.POST("/api/register", user.CreateUserHandler(user.GetUserByUsername(mongodb), user.CreateUser(mongodb)))
	e.GET("/api/user/:username", user.GetUserHandler(user.GetUserByUsername(mongodb)))

	u := e.Group("")
	config := middleware.JWTConfig{
		Claims:     &login.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	u.Use(middleware.JWTWithConfig(config))
	u.GET("/api/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "admin&user can see")
	}).Name = "users"

	a := u.Group("/api", mw.OnlyAdmin)
	a.GET("/admin", func(c echo.Context) error {
		return c.String(http.StatusOK, "only admin can see")
	}).Name = "users"

	// Start server
	go func() {
		if err := e.Start(":8000"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
