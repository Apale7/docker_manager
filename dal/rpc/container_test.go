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
	for _, v := range resp {
		// continue
		fmt.Printf("%+v\n", v)
	}
}

func TestGetContainer(t *testing.T) {
	_, err := GetContainer(ctx, "93177708cacd7147c7ab30ef32c903c0678970baa1a068e35ad3a52908561db3")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestAll(t *testing.T) {
	containers, err := GetAllContainers(ctx)
	if err != nil {
		t.Logf("GetAllContainers error: %+v", err)
		t.FailNow()
	}
	n := len(containers)
	images, err := GetAllImages(ctx)
	if err != nil {
		t.Logf("GetAllImages error: %+v", err)
		t.FailNow()
	}
	image := images[len(images)-1]
	containerID, err := CreateContainer(ctx, image.Id, image.Author, "cyq_test")
	if err != nil {
		t.Logf("CreateContainer error: %+v", err)
		// t.FailNow()
	}
	containers, err = GetAllContainers(ctx)
	if err != nil || len(containers) != n+1 {
		t.Logf("CreateContainer error: %+v", err)
		// t.FailNow()
	}
	n = len(containers)
	err = DeleteContainer(ctx, containerID)
	if err != nil {
		t.Logf("DeleteContainer error: %+v", err)
		t.FailNow()
	}
	containers, err = GetAllContainers(ctx)
	if err != nil || len(containers) != n-1 {
		t.Logf("DeleteContainer error: %+v", err)
		t.FailNow()
	}
}

func TestDeleteContainer(t *testing.T) {
	err := DeleteContainer(ctx, "1d240c395a70ef641c41923b375eaad93c174c112fe991935533cde4933bc7dd")
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
