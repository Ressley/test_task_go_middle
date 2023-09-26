package requests

type PostMessageRequest struct {
	Destination *string `json:"destination" validate:"required"`
	Data        []byte  `json:"data" validate:"required"`
}
