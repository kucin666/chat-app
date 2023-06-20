package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thxrhmn/chat-app/handlers"
	"github.com/thxrhmn/chat-app/pkg/middleware"
	"github.com/thxrhmn/chat-app/pkg/mysql"
	"github.com/thxrhmn/chat-app/repositories"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/user/:id", h.GetUser)
	e.GET("/users", h.FindUsers)
	e.DELETE("/user/:id", middleware.Auth(h.DeleteUser))
}
