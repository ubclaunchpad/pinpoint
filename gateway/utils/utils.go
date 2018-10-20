package utils

import (
	"context"
	"os"

	"google.golang.org/grpc/metadata"
)

// FirstString returns the first element from an array of strings
func FirstString(strs []string) string {
	if strs != nil && len(strs) > 0 {
		return strs[0]
	}
	return ""
}

// SecureContext takes in a context and then adds in authentication token to it
func SecureContext(ctx context.Context) context.Context {
	// set up ctx for future communication
	md := metadata.Pairs("token", os.Getenv("PINPOINT_CORE_TOKEN"))
	ctxnew := metadata.NewOutgoingContext(ctx, md)

	return ctxnew
}
