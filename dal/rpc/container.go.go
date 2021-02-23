package rpc

import (
	"context"
	containerManager "docker_manager/proto/container_server"
)

func GetAllContainers(ctx context.Context) (containers []*containerManager.ContainerAttr, err error) {
	resp, err := containerManagerClient.ListContainers(ctx, &containerManager.ListContainers_Request{})
	if err != nil {
		return
	}
	return resp.Containers, nil
}

func GetContainer(ctx context.Context, containerID string) (container *containerManager.ContainerAttr, err error) {
	resp, err := containerManagerClient.GetContainer(ctx, &containerManager.GetContainer_Request{ContainerId: containerID})
	if err != nil {
		return
	}

	return resp.ContainerAttr, nil
}
