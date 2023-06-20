package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	dto "github.com/thxrhmn/chat-app/dto/result"
	roomdto "github.com/thxrhmn/chat-app/dto/room"
	"github.com/thxrhmn/chat-app/models"
	"github.com/thxrhmn/chat-app/repositories"
)

type handlerRoom struct {
	RoomRepository repositories.RoomRepository
}

func HandlerRoom(RoomRepository repositories.RoomRepository) *handlerRoom {
	return &handlerRoom{RoomRepository}
}

func (h *handlerRoom) CreateRoom(c echo.Context) error {
	request := roomdto.RoomRequest{
		Name: c.FormValue("name"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	room := models.Room{
		Name:      request.Name,
		CreatedBy: int(userId),
	}

	room, err = h.RoomRepository.CreateRoom(room)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	room, _ = h.RoomRepository.GetRoom(room.ID)
	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: room})
}

func (h *handlerRoom) GetRoom(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	room, err := h.RoomRepository.GetRoom(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: room})
}

func (h *handlerRoom) FindRooms(c echo.Context) error {
	rooms, err := h.RoomRepository.FindRooms()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: rooms})
}

func (h *handlerRoom) UpdateRoom(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := roomdto.RoomRequest{
		Name: c.FormValue("name"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	room, _ := h.RoomRepository.GetRoom(id)

	if request.Name != "" {
		room.Name = request.Name
	}

	data, err := h.RoomRepository.UpdateRoom(room)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}

func (h *handlerRoom) DeleteRoom(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	room, err := h.RoomRepository.GetRoom(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	data, err := h.RoomRepository.DeleteRoom(room, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: http.StatusOK, Data: data})
}
