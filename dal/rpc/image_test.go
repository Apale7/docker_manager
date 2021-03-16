package rpc

import (
	"context"
	"docker_manager/dal/db"
	"docker_manager/dto"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetImages(t *testing.T) {
	ctx := context.Background()
	images, err := GetAllImages(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, i := range images {
		err := db.CreateImage(ctx, 2, dto.RPCImageToModelImage(i))
		if err != nil {
			logrus.Errorln(err)
			t.FailNow()
		}
		fmt.Println(i.Id + " created.")
	}
}
