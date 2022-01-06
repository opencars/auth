package model

var (
	ErrBlacklistRecordNotFound = NewError("blacklist.record.not_found")
	ErrTokenNotFound           = NewError("token.not_found")

	ErrInvalidLimit  = NewError("query.invalid_limit")
	ErrInvalidOffset = NewError("query.invalid_offset")

	ErrInvalidToken = NewError("auth.token.is_not_valid")
	ErrTokenRevoked = NewError("auth.token.revoked")
	ErrAccessDenied = NewError("auth.access_denied")
)

type Error struct {
	text string
}

func NewError(text string) Error {
	return Error{
		text: text,
	}
}

func (e Error) Error() string {
	return e.text
}
