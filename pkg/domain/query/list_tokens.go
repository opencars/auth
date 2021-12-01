package query

import (
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type ListTokens struct {
	UserID string
	Limit  string
	Offset string
}

func (q *ListTokens) GetLimit() uint64 {
	if q.Limit == "" {
		return 10
	}

	num, err := strconv.ParseInt(q.Limit, 10, 64)
	if err != nil {
		panic(err)
	}

	if num < 0 {
		return 10
	}

	return uint64(num)
}

func (q *ListTokens) GetOffset() uint64 {
	if q.Offset == "" {
		return 0
	}

	num, err := strconv.ParseInt(q.Offset, 10, 64)
	if err != nil {
		panic(err)
	}

	if num < 0 {
		return 10
	}

	return uint64(num)
}

func (q *ListTokens) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.UserID,
			validation.Required.Error("required"),
		),
		validation.Field(
			&q.Limit,
			is.Int.Error("is_not_integer"),
		),
		validation.Field(
			&q.Offset,
			is.Int.Error("is_not_integer"),
		),
	)
}
