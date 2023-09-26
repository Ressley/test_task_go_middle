package requests

type PostChallengeRequest struct {
	Destination *string `json:"destination" validate:"required"`
}
