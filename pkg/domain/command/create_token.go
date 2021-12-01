package command

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateToken struct {
	UserID string
	Name   string
}

func (c *CreateToken) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.UserID,
			validation.Required.Error("required"),
		),
		validation.Field(
			&c.Name,
			validation.Required.Error("required"),
			validation.Length(2, 64).Error("invalid_size"),
		),
	)
}
