package roomdto

type RoomRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}
