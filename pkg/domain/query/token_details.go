package query

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type TokenDetails struct {
	UserID  string
	TokenID string
}

func (q *TokenDetails) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.UserID,
			validation.Required.Error("required"),
		),
		validation.Field(
			&q.TokenID,
			validation.Required.Error("required"),
		),
	)
}
