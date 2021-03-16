package handler

import (
	"context"
	"docker_manager/dal/db"
	"docker_manager/dal/rpc"
	"docker_manager/dto"
	"docker_manager/proto/base"
	"docker_manager/proto/docker_manager"

	"github.com/Apale7/common/utils"
	"github.com/sirupsen/logrus"
)

// CreateContainer 先创建container，再记录user与container的关系
func CreateContainer(ctx context.Context, req *docker_manager.CreateContainerRequest) (resp *docker_manager.CreateContainerResponse, err error) {
	resp = &docker_manager.CreateContainerResponse{}
	resp.BaseResp = &base.BaseResp{}
	containerAttr, err := rpc.CreateContainer(ctx, req.ImageId, req.Username, req.ContainerName)

	if err != nil {
		logrus.Warnf("CreateContainer error: %v", err)
		// resp.BaseResp.
		return
	}
	err = db.CreateContainer(ctx, req.UserId, containerAttr.ToDBContainer())
	if err != nil {
		logrus.Warnf("CreateContainer error: %v", err)
		return
	}

	return
}

// DeleteContainer 先删除container，再删除user与container的关系
func DeleteContainer(ctx context.Context, req *docker_manager.DeleteContainerRequest) (resp *docker_manager.DeleteContainerResponse, err error) {
	resp = &docker_manager.DeleteContainerResponse{}
	err = db.DeleteContainer(ctx, req.UserId, req.ContainerId)
	if err != nil {
		logrus.Warnf("DeleteContainer error: %v", err)
	}
	go utils.ProtectRun(func() {
		err = rpc.DeleteContainer(ctx, req.ContainerId)
		if err != nil { //失败了不处理，// todo: 自动删除长期不使用的容器
			logrus.Warnf("DeleteContainer error: %v", err)
			return
		}
	})

	return
}

func GetContainer(ctx context.Context, req *docker_manager.GetContainerRequest) (resp *docker_manager.GetContainerResponse, err error) {
	resp = &docker_manager.GetContainerResponse{}
	containers, err := db.GetContainer(ctx, req.UserId, req.ContainerId)

	if err != nil {
		return
	}

	resp.Containers = make([]*docker_manager.Container, 0, len(containers))
	for _, c := range containers {
		resp.Containers = append(resp.Containers, dto.ModelContainerToDockerManagerContainer(c))
	}
	return
}
