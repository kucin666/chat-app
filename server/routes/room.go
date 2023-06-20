package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thxrhmn/chat-app/handlers"
	"github.com/thxrhmn/chat-app/pkg/middleware"
	"github.com/thxrhmn/chat-app/pkg/mysql"
	"github.com/thxrhmn/chat-app/repositories"
)

func RoomRoutes(e *echo.Group) {
	roomRepository := repositories.RepositoryRoom(mysql.DB)
	h := handlers.HandlerRoom(roomRepository)

	e.POST("/room", middleware.Auth(h.CreateRoom))
	e.GET("/room/:id", h.GetRoom)
	e.GET("/rooms", h.FindRooms)
	e.PATCH("/room/:id", middleware.Auth(h.UpdateRoom))
	e.DELETE("/room/:id", middleware.Auth(h.DeleteRoom))
}
