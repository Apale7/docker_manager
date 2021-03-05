package handler

import (
	"context"
	"docker_manager/dal/rpc"
	"docker_manager/dto"
	"docker_manager/proto/docker_manager"
)

func GetAllContainers(ctx context.Context) (resp *docker_manager.GetAllContainersResponse, err error) {
	containerAttrs, err := rpc.GetAllContainers(ctx)
	if err != nil {
		return
	}
	containers := make([]*docker_manager.Container, len(containerAttrs))
	for i := range containerAttrs {
		containers[i] = dto.ToContainer(containerAttrs[i])
	}
	resp = &docker_manager.GetAllContainersResponse{Containers: containers}

	return
}

func GetContainers() {

}
