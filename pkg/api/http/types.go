package http

import "context"

const (
	HeaderTokenID   = "X-Auth-Id"
	HeaderTokenName = "X-Auth-Name"
	HeaderUserID    = "X-User-Id"
)

type HeaderKey int

const (
	UserIDHeaderKey HeaderKey = iota
)

type CreateTokenRequest struct {
	Name string `json:"name"`
}

func UserIDFromContext(ctx context.Context) string {
	return ctx.Value(UserIDHeaderKey).(string)
}

func WithUserID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, UserIDHeaderKey, id)
}
