package room

type RoomRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}
