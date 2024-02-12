package test

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"
)

// NewController starts a new controller for the mock
func NewController(t *testing.T) (*gomock.Controller, context.Context) {
	ctx := context.Background()
	return gomock.WithContext(ctx, t)
}
