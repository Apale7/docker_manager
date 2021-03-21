package rpc

import (
	"context"
	Manage "docker_manager/proto/container_server"
	"fmt"
	"testing"
)

func TestGetImages(t *testing.T) {
	ctx := context.Background()
	images, err := GetAllImages(ctx)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, i := range images {
		// err := db.CreateImage(ctx, 2, dto.RPCImageToModelImage(i))
		// if err != nil {
		// 	logrus.Errorln(err)
		// 	t.FailNow()
		// }
		fmt.Printf("%+v\n", i)
	}
}

func TestCreateImage(t *testing.T) {
	imageAttr, err := PullImage(ctx, "ubuntu", "14.04", &Manage.AuthConfig{Username: "123", Password: "456"})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Printf("%+v", imageAttr)
}

func TestDeleteImage(t *testing.T) {
	err := DeleteImage(ctx, "sha256:df043b4f0cf196749a9a426080f433b76cabf6b37dde2edefef317ba54c713c7", true)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
