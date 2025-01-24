package validatex

import "context"

func GetValFromCtx(ctx context.Context, key string) string {
	val := ctx.Value(key)
	if val == nil {
		return ""
	}
	return val.(string)
}
