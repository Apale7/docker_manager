package handler

import (
	"context"
	"docker_manager/dal/db"
	"docker_manager/dal/db/model"
	"docker_manager/dal/rpc"
	"docker_manager/dto"
	"docker_manager/proto/base"
	"docker_manager/proto/docker_manager"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

// CreateContainer 先创建container，再记录user与container的关系
func CreateContainer(ctx context.Context, req *docker_manager.CreateContainerRequest) (resp *docker_manager.CreateContainerResponse, err error) {
	resp = &docker_manager.CreateContainerResponse{}
	resp.BaseResp = &base.BaseResp{}
	containerAttr, err := rpc.CreateContainer(ctx, req.ImageId, req.Username, req.ContainerName)

	if err != nil {
		logrus.Warnf("RPC CreateContainer error: %v", err)
		// resp.BaseResp.
		return
	}
	container := containerAttr.ToDBContainer()
	container.Name = req.ContainerName
	err = db.CreateContainer(ctx, req.UserId, container)
	if err != nil {
		logrus.Warnf("DB CreateContainer error: %v", err)
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
	err = rpc.DeleteContainer(ctx, req.ContainerId)
	if err != nil { //失败了不处理，// todo: 自动删除长期不使用的容器
		logrus.Warnf("DeleteContainer error: %v", err)
		return
	}

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

func StartContainer(ctx context.Context, req *docker_manager.StartContainerRequest) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}
	_, err = db.GetUserContainer(ctx, req.UserId, req.ContainerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return resp, errors.WithStack(errors.New("无权限启动该容器"))
	}
	if err != nil {
		return resp, errors.WithStack(errors.New("服务器错误"))
	}
	err = rpc.StartContainer(ctx, req.ContainerId)
	if err != nil {
		return resp, errors.WithStack(errors.New("启动容器失败"))
	}
	err = db.UpdateContainer(ctx, &model.Container{ContainerID: req.ContainerId, Status: int8(docker_manager.Container_Running)})
	return
}
func StopContainer(ctx context.Context, req *docker_manager.StopContainerRequest) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}
	_, err = db.GetUserContainer(ctx, req.UserId, req.ContainerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return resp, errors.WithStack(errors.New("无权限关闭该容器"))
	}
	if err != nil {
		return resp, errors.WithStack(errors.New("服务器错误"))
	}
	err = rpc.StopContainer(ctx, req.ContainerId)
	if err != nil {
		return resp, errors.WithStack(errors.New("关闭容器失败"))
	}
	err = db.UpdateContainer(ctx, &model.Container{ContainerID: req.ContainerId, Status: int8(docker_manager.Container_Paused)})
	return
}
func RestartContainer(ctx context.Context, req *docker_manager.RestartContainerRequest) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}
	err = db.UpdateContainer(ctx, &model.Container{ContainerID: req.ContainerId, Status: int8(docker_manager.Container_Restarting)})
	if err != nil {
		return resp, errors.WithStack(errors.New("服务器错误"))
	}
	_, err = db.GetUserContainer(ctx, req.UserId, req.ContainerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return resp, errors.WithStack(errors.New("无权限重启该容器"))
	}
	if err != nil {
		return resp, errors.WithStack(errors.New("服务器错误"))
	}
	err = rpc.RestartContainer(ctx, req.ContainerId)
	if err != nil {
		return resp, errors.WithStack(errors.New("重启容器失败"))
	}
	err = db.UpdateContainer(ctx, &model.Container{ContainerID: req.ContainerId, Status: int8(docker_manager.Container_Running)})
	if err != nil {
		return resp, errors.WithStack(errors.New("服务器错误"))
	}
	return
}
