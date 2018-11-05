package service

import (
	"context"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

func authUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "missing context metadata")
	}
	if err := validate(meta, os.Getenv("PINPOINT_CORE_TOKEN")); err != nil {
		return nil, err
	}

	// continue
	return handler(ctx, req)
}

func authStreamingInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	meta, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "missing context metadata")
	}
	if err := validate(meta, os.Getenv("PINPOINT_CORE_TOKEN")); err != nil {
		return err
	}

	// continue
	return handler(srv, stream)
}

// validate checks given metadata for key and returns an appropriate gRPC error
// if anything goes wrong
func validate(meta metadata.MD, key string) error {
	keys, ok := meta["key"]
	if !ok || len(meta["key"]) == 0 {
		return grpc.Errorf(codes.Unauthenticated, "no key provided")
	}
	if keys[0] != key {
		return grpc.Errorf(codes.Unauthenticated, "invalid key")
	}
	return nil
}
