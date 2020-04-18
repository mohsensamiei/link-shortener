package auth

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const authToken = "authToken"

func AddTokenToContext(ctx context.Context, token string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(authToken, token))
}
func GetTokenFromContext(ctx context.Context) (string, bool) {
	pair, ok := metadata.FromIncomingContext(ctx)
	if ok == false {
		return "", false
	}
	return pair.Get(authToken)[0], true
}
