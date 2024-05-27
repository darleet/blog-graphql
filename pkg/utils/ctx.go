package utils

import (
	"context"
)

type ContextKey struct{}

func GetUserID(ctx context.Context) string {
	userID := ctx.Value(ContextKey{})
	if userID == nil {
		return ""
	}
	return userID.(string)
}

func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ContextKey{}, userID)
}
