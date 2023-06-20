package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thxrhmn/chat-app/handlers"
	"github.com/thxrhmn/chat-app/pkg/middleware"
	"github.com/thxrhmn/chat-app/pkg/mysql"
	"github.com/thxrhmn/chat-app/repositories"
)

func ChatRoutes(e *echo.Group) {
	chatRepository := repositories.RepositoryChat(mysql.DB)
	h := handlers.HandlerChat(chatRepository)

	e.POST("/chat", middleware.Auth(h.CreateChat))
	e.GET("/chat/:id", h.GetChat)
	e.GET("/chats", h.FindChats)
	e.GET("/room/:roomID/chat/:chatID", h.GetChatByRoomID)
	e.GET("/room/:roomID/chats", h.FindChatsByRoomID)
	e.PATCH("/chat/:id", middleware.Auth(h.UpdateChat))
	e.DELETE("/chat/:id", middleware.Auth(h.DeleteChat))
}
