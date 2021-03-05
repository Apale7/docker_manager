package dto

import (
	containerManager "docker_manager/proto/container_server"
	"docker_manager/proto/docker_manager"
	"time"
)

func ToContainer(m *containerManager.ContainerAttr) *docker_manager.Container {
	ti, _ := time.Parse(time.RFC3339Nano, m.Created)
	return &docker_manager.Container{
		Id:      m.Id,
		Status:  docker_manager.Container_ContainerStatus(m.Status),
		Created: ti.Unix(),
		Image:   m.Image,
		Name:    m.Name,
	}
}
