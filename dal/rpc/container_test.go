package rpc

import (
	"context"
	"fmt"
	"testing"
)

var ctx = context.Background()

func TestGetContainers(t *testing.T) {
	resp, err := GetAllContainers(ctx)
	if err != nil {
		t.FailNow()
	}
	for _, i := range resp {
		fmt.Println(i.Name, i.Status.String())
	}
}

func TestGetContainer(t *testing.T) {
	_, err := GetContainer(ctx, "93177708cacd7147c7ab30ef32c903c0678970baa1a068e35ad3a52908561db3")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestDeleteContainer(t *testing.T) {
	err := DeleteContainer(ctx, "422a9705bc9c2f678df1be46158be2c7e770a3b94b9533bf6796b71a32cc226d")
	if err != nil {
		t.Logf("DeleteContainer error: %+v", err)
		t.FailNow()
	}
}

func TestPruneContainer(t *testing.T) {
	err := PruneContainers(ctx)
	if err != nil {
		t.Logf("PruneContainer error: %+v", err)
		t.FailNow()
	}
	TestGetContainers(t)
}

func TestStopContainers(t *testing.T) {
	c, _ := GetAllContainers(ctx)

	err := StopContainer(ctx, c[1].Id)
	if err != nil {
		t.Logf("StopContainer error: %+v", err)
		t.FailNow()
	}

	err = StartContainer(ctx, c[1].Id)
	if err != nil {
		t.Logf("StartContainer error, err: %+v", err)
	}

}

func TestStartContainer(t *testing.T) {
	err := StartContainer(ctx, "fd5354ec483f56565746aec4c642bc418115c162f2c1bc18751004a93a7a5686")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
