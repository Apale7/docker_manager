package main

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	"docker_manager/dal/rpc"
	"docker_manager/handler"
	"docker_manager/proto/docker_manager"
)

// DockerManagerServer DockerManagerServer impl
type DockerManagerServer struct {
}

// CreateContainer create a container for a user
func (DockerManagerServer) CreateContainer(ctx context.Context, req *docker_manager.CreateContainerRequest) (resp *docker_manager.CreateContainerResponse, err error) {
	fmt.Println("CreateContainer called.")
	return handler.CreateContainer(ctx, req)
}

// DeleteContainer delete a container for a user
func (DockerManagerServer) DeleteContainer(ctx context.Context, req *docker_manager.DeleteContainerRequest) (resp *docker_manager.DeleteContainerResponse, err error) {
	fmt.Println("DeleteContainer called.")
	return handler.DeleteContainer(ctx, req)
}

// GetContainer get containers by container_id or user_id
func (DockerManagerServer) GetContainer(ctx context.Context, req *docker_manager.GetContainerRequest) (resp *docker_manager.GetContainerResponse, err error) {
	fmt.Println("GetContainer called.")
	return handler.GetContainer(ctx, req)
}

// PruneContainers delete all unused containers
func (DockerManagerServer) PruneContainers(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	fmt.Println("PruneContainers called.")
	return &emptypb.Empty{}, rpc.PruneContainers(ctx)
}

// CreateImage create a image for a user
func (DockerManagerServer) CreateImage(ctx context.Context, req *docker_manager.CreateImageRequest) (resp *docker_manager.CreateImageResponse, err error) {
	fmt.Println("CreateImage called.")
	return handler.CreateImage(ctx, req)
}

// DeleteImage delete a image for a user
func (DockerManagerServer) DeleteImage(ctx context.Context, req *docker_manager.DeleteImageRequest) (resp *docker_manager.DeleteImageResponse, err error) {
	fmt.Println("DeleteImage called.")
	return handler.DeleteImage(ctx, req)
}

// GetImage get images by image_id or user_id
func (DockerManagerServer) GetImage(ctx context.Context, req *docker_manager.GetImageRequest) (resp *docker_manager.GetImageResponse, err error) {
	fmt.Println("GetImage called.")
	return handler.GetImage(ctx, req)
}

// PruneImages delete all unused images
func (DockerManagerServer) PruneImages(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	fmt.Println("PruneImages called.")
	return nil, rpc.PruneImages(ctx)
}

func (DockerManagerServer) StartContainer(ctx context.Context, req *docker_manager.StartContainerRequest) (*emptypb.Empty, error) {
	fmt.Println("StartContainer called.")
	return handler.StartContainer(ctx, req)
}
func (DockerManagerServer) StopContainer(ctx context.Context, req *docker_manager.StopContainerRequest) (*emptypb.Empty, error) {
	fmt.Println("StopContainer called.")
	return handler.StopContainer(ctx, req)
}
func (DockerManagerServer) RestartContainer(ctx context.Context, req *docker_manager.RestartContainerRequest) (*emptypb.Empty, error) {
	fmt.Println("RestartContainer called.")
	return handler.RestartContainer(ctx, req)
}
