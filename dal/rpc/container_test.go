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
		// err := db.CreateContainer(ctx, 2, dto.RPCContainerToModelContainer(v))
		// if err != nil {
		// 	logrus.Errorln("name too long: " + v.Name)
		// }
		fmt.Printf("%+v %s\n", v.Id, v.Status.String())
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
