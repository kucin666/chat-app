package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thxrhmn/chat-app/handlers"
	"github.com/thxrhmn/chat-app/pkg/middleware"
	"github.com/thxrhmn/chat-app/pkg/mysql"
	"github.com/thxrhmn/chat-app/repositories"
)

func AuthRoutes(e *echo.Group) {
	authRepository := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authRepository)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
}
