package models_bot

type StartRequestModel struct {
	Symbol string `json:"symbol" validate:"required,alphanum,uppercase"`
	IsMass bool   `json:"isMass" validate:"boolean"`
}
