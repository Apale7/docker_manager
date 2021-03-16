package dto

import (
	"docker_manager/dal/db/model"
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
