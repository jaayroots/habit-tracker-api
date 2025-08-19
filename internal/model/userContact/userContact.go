package model

import (
	"github.com/google/uuid"
	"github.com/habit-tracker-api/enums"
)

type (
	UserContact struct {
		ID          uuid.UUID         `json:"id"`
		ContactType enums.ContactType `json:"contact_type"`
		Value       string            `json:"value"`
	}

	UserContactReq struct {
		ID          uuid.UUID `json:"id"`
		ContactType string    `json:"contact_type" validate:"required"`
		Value       string    `json:"value" validate:"required"`
		IsDelete    bool      `json:"is_delete"`
	}

	UserContactRes struct {
		ID          uuid.UUID `json:"id"`
		ContactType string    `json:"contact_type"`
		Value       string    `json:"Value"`
	}
)
