package rpc

import (
	"context"
	containerManager "docker_manager/proto/container_server"
	"fmt"

	"github.com/satori/go.uuid"
)

func GetAllContainers(ctx context.Context) (containers []*containerManager.ContainerAttr, err error) {
	resp, err := containerManagerClient.ListContainers(ctx, &containerManager.ListContainers_Request{})
	if err != nil {
		return
	}
	return resp.Containers, nil
}

func PruneContainers(ctx context.Context) (err error) {
	_, err = containerManagerClient.PruneContainers(ctx, &containerManager.PruneContainers_Request{})

	return
}

func GetContainer(ctx context.Context, containerID string) (container *containerManager.ContainerAttr, err error) {
	resp, err := containerManagerClient.GetContainer(ctx, &containerManager.GetContainer_Request{ContainerId: containerID})
	if err != nil {
		return
	}

	return resp.ContainerAttr, nil
}

func DeleteUnusedContainer(ctx context.Context) (err error) {
	_, err = containerManagerClient.PruneContainers(ctx, &containerManager.PruneContainers_Request{})

	return
}

//todo start前get一下，如果已经打开则不start
func StartContainer(ctx context.Context, containerID string) (err error) {
	_, err = containerManagerClient.StartContainer(ctx, &containerManager.StartContainer_Request{ContainerId: containerID})

	return
}

func StopContainer(ctx context.Context, containerID string) (err error) {
	_, err = containerManagerClient.StopContainer(ctx, &containerManager.StopContainer_Request{ContainerId: containerID})

	return
}

func RestartContainer(ctx context.Context, containerID string) (err error) {
	_, err = containerManagerClient.RestartContainer(ctx, &containerManager.RestartContainer_Request{ContainerId: containerID})

	return
}

func DeleteContainer(ctx context.Context, containerID string) (err error) {
	_, err = containerManagerClient.RemoveContainer(ctx, &containerManager.RemoveContainer_Request{ContainerId: containerID, Force: true})

	return
}

func CreateContainer(ctx context.Context, imageID, username, containerName string) (containerAttr *containerManager.ContainerAttr, err error) {
	req := &containerManager.CreateContainer_Request{
		ImageId:       imageID,
		Username:      username,
		ContainerName: fmt.Sprintf("%s_%s", username, uuid.NewV4().String()), //业务中不关心k8s的container_name
	}
	resp, err := containerManagerClient.CreateContainer(ctx, req)

	if err != nil {
		return
	}

	return resp.ContainerAttr, nil
}
