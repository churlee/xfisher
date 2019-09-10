package dto

type CreateFeedbackDto struct {
	Content string `json:"content" validate:"required"`
}
