package rpc

import (
	"context"
	"testing"
)

func TestGetContainers(t *testing.T) {
	ctx := context.Background()
	resp, err := GetContainers(ctx)
	if err != nil {
		t.FailNow()
	}
	t.Logf("%+v", resp)
}
