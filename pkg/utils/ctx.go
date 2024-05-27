package utils

import (
	"context"
	"fmt"
)

type ContextKey struct{}

func GetUserID(ctx context.Context) (string, error) {
	userID := ctx.Value(ContextKey{})
	if userID == nil {
		return "", fmt.Errorf("userID not found in context")
	}
	return userID.(string), nil
}
