package model

import "docker_manager/proto/docker_manager"

func (c *Container) ToGRPCContainer() *docker_manager.Container {
	return &docker_manager.Container{
		Id:      c.ContainerID,
		Created: c.CreatedAt.Local().Unix(),
		Status:  docker_manager.Container_ContainerStatus(c.Status),
		Image:   c.ImageID,
		Name:    c.Name,
	}
}
