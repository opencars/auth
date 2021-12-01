package command

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ResetToken struct {
	UserID  string
	TokenID string
}

func (c *ResetToken) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.UserID,
			validation.Required.Error("required"),
		),
		validation.Field(
			&c.TokenID,
			validation.Required.Error("required"),
		),
	)
}
