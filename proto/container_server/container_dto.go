package Manage

import (
	"docker_manager/dal/db/model"
)

func (c *ContainerAttr) ToDBContainer() *model.Container {
	return &model.Container{
		ContainerID: c.Id,
		Status:      int8(c.Status),
		ImageID:     c.Image,
		Name:        c.Name,
	}
}
