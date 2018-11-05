package service

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func newAuthInterceptors(token string) (
	unaryInterceptor grpc.UnaryServerInterceptor,
	streamInterceptor grpc.StreamServerInterceptor,
) {
	unaryInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, grpc.Errorf(codes.Unauthenticated, "missing context metadata")
		}
		if err := validate(meta, token); err != nil {
			return nil, err
		}

		// continue
		return handler(ctx, req)
	}

	streamInterceptor = func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		meta, ok := metadata.FromIncomingContext(stream.Context())
		if !ok {
			return grpc.Errorf(codes.Unauthenticated, "missing context metadata")
		}
		if err := validate(meta, token); err != nil {
			return err
		}

		// continue
		return handler(srv, stream)
	}
	return
}

// validate checks given metadata for key and returns an appropriate gRPC error
// if anything goes wrong
func validate(meta metadata.MD, key string) error {
	keys, ok := meta["authorization"]
	if !ok || len(meta["authorization"]) == 0 {
		return grpc.Errorf(codes.Unauthenticated, "no key provided")
	}
	if keys[0] != key {
		return grpc.Errorf(codes.Unauthenticated, "invalid key")
	}
	return nil
}
