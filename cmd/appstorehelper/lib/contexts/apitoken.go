package contexts

import (
	"context"
	"fmt"
)

var (
	APITokenContextKey = contextKey{"APIToken"}
)

func ContextWithAPIToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, APITokenContextKey, token)
}

func APITokenFromContext(ctx context.Context) (string, error) {
	token, ok := ctx.Value(APITokenContextKey).(string)
	if !ok {
		return "", NewNoValueError(fmt.Errorf("key: %v", APITokenContextKey))
	}

	return token, nil
}
