package rpc

import (
	"context"
	containerManager "docker_manager/proto/container_server"
)

func GetImages(ctx context.Context) (images []*containerManager.ImageAttr, err error) {
	resp, err := containerManagerClient.ListImages(ctx, &containerManager.ListImages_Request{})
	if err != nil {
		return
	}
	return resp.Images, nil
}
