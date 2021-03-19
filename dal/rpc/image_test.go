package rpc

import (
	"context"
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
		if i.Id == "sha256:852698725025535f9185f477e385577db6478d022593d662707b332bb0553d56" {
			fmt.Println(i.Id + " exist.")
		}
	}
}
