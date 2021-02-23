package rpc

import (
	"context"
	containerManager "docker_manager/proto/container_server"
)

func GetContainers(ctx context.Context) (containers []*containerManager.ContainerAttr, err error) {
	resp, err := containerManagerClient.ListContainers(ctx, &containerManager.ListContainers_Request{})
	if err != nil {
		return
	}
	return resp.Containers, nil
}

