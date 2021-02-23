package rpc

import (
	"context"
	"fmt"
	"testing"
)

var ctx = context.Background()

func TestGetContainers(t *testing.T) {
	_, err := GetAllContainers(ctx)
	if err != nil {
		t.FailNow()
	}
	// fmt.Printf("%+v", resp)
}

func TestGetContainer(t *testing.T) {
	_, err := GetContainer(ctx, "93177708cacd7147c7ab30ef32c903c0678970baa1a068e35ad3a52908561db3")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
