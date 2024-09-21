package context

import "golang.org/x/net/context"

type key string

const (
	timeoutKey  key = "TimeoutKey"
	deadlineKey key = "DeadlineKey"
)

func Setup(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, timeoutKey, "timeout exceeded")
	ctx = context.WithValue(ctx, deadlineKey, "deadline exceeded")
	return ctx
}

func GetValue(ctx context.Context, k key) string {
	if val, ok := ctx.Value(k).(string); ok {
		return val
	}
	return ""
}
