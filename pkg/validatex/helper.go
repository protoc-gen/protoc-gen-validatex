package validatex

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func GetValFromCtx(ctx context.Context, key string) string {
	s := getFromGrpc(ctx, key)
	if s != "" {
		return s
	}

	return getFromCtx(ctx, key)
}

func getFromGrpc(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		vals := md.Get(key)
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return ""
}

func getFromCtx(ctx context.Context, key string) string {
	val := ctx.Value(key)
	if val == nil {
		return ""
	}
	return val.(string)
}
