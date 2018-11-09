package utils

import (
	"context"
	"testing"
)

func TestCredentials_GetRequestMetadata(t *testing.T) {
	token := "hello world"
	c := NewCredentials(token)
	got, err := c.GetRequestMetadata(context.Background())
	if err != nil {
		t.Error(err)
	}
	if got["authorization"] != token {
		t.Errorf("expected %s, got %s", token, got["authorization"])
	}

	// get coverage on simple getter
	c.RequireTransportSecurity()
}
