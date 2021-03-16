package dto

import (
	"docker_manager/dal/db/model"
	containerManager "docker_manager/proto/container_server"
	"docker_manager/proto/docker_manager"
)

func ModelContainerToDockerManagerContainer(c *model.Container) *docker_manager.Container {
	return &docker_manager.Container{
		Id:      c.ContainerID,
		Created: c.CreatedAt.Local().Unix(),
		Status:  docker_manager.Container_ContainerStatus(c.Status),
		Image:   c.ImageID,
		Name:    c.Name,
	}
}

func RPCContainerToModelContainer(c *containerManager.ContainerAttr) *model.Container {
	return &model.Container{
		ContainerID: c.Id,
		Status:      int8(c.Status),
		ImageID:     c.Image,
		Name:        c.Name,
	}
}
