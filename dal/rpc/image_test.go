package rpc

import (
	"context"
	"testing"
)

func TestGetImages(t *testing.T) {
	ctx := context.Background()
	_, err := GetImages(ctx)
	if err != nil {
		t.FailNow()
	}
	// fmt.Printf("%+v", resp)
}
