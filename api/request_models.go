package api

import "github.com/google/uuid"

type codeRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=9"`
}

type codeValidation struct {
	Id             uuid.UUID `json:"id" validate:"required,uuid4"`
	ValidationCode int       `json:"validation_code" validate:"required"`
}
