package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	chatdto "github.com/thxrhmn/chat-app/dto/chat"
	dto "github.com/thxrhmn/chat-app/dto/result"
	"github.com/thxrhmn/chat-app/models"
	"github.com/thxrhmn/chat-app/repositories"
)

type handlerChat struct {
	ChatRepository repositories.ChatRepository
}

func HandlerChat(ChatRepository repositories.ChatRepository) *handlerChat {
	return &handlerChat{ChatRepository}
}

func (h *handlerChat) CreateChat(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	roomID, _ := strconv.Atoi(c.FormValue("room_id"))

	request := chatdto.ChatRequest{
		Title:  c.FormValue("title"),
		RoomID: roomID,
		UserID: int(userId),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	chat := models.Chat{
		Title:  request.Title,
		RoomID: request.RoomID,
		UserID: int(userId),
	}

	chat, err = h.ChatRepository.CreateChat(chat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	chat, _ = h.ChatRepository.GetChat(chat.ID)
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: chat})
}

func (h *handlerChat) GetChat(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	chat, err := h.ChatRepository.GetChat(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: chat})
}

func (h *handlerChat) FindChats(c echo.Context) error {
	chats, err := h.ChatRepository.FindChats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: chats})
}

func (h *handlerChat) FindChatsByRoomID(c echo.Context) error {
	roomID, _ := strconv.Atoi(c.Param("roomID"))

	chats, err := h.ChatRepository.FindChatsByRoomID(roomID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: chats})
}

func (h *handlerChat) GetChatByRoomID(c echo.Context) error {
	roomID, _ := strconv.Atoi(c.Param("roomID"))
	chatID, _ := strconv.Atoi(c.Param("chatID"))

	chat, err := h.ChatRepository.GetChatByRoomID(roomID, chatID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: chat})
}

func (h *handlerChat) UpdateChat(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	id, _ := strconv.Atoi(c.Param("id"))

	roomID, _ := strconv.Atoi(c.Param("room_id"))

	request := chatdto.ChatRequest{
		Title:  c.FormValue("title"),
		UserID: int(userId),
		RoomID: roomID,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	chat, _ := h.ChatRepository.GetChat(id)

	if request.Title != "" {
		chat.Title = request.Title
	}

	if request.RoomID != 0 {
		chat.RoomID = request.RoomID
	}

	data, err := h.ChatRepository.UpdateChat(chat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerChat) DeleteChat(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	chat, err := h.ChatRepository.GetChat(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	data, err := h.ChatRepository.DeleteChat(chat, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}
