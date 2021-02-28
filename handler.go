package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"docker_manager/proto/docker_manager"
)

type DockerManagerServer struct {
}

func (DockerManagerServer) CreateContainer(context.Context, *docker_manager.CreateContainerRequest) (*docker_manager.CreateContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContainer not implemented")
}
func (DockerManagerServer) DeleteContainer(context.Context, *docker_manager.DeleteContainerRequest) (*docker_manager.DeleteContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContainer not implemented")
}
func (DockerManagerServer) GetContainer(context.Context, *docker_manager.GetContainerRequest) (*docker_manager.GetContainerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainer not implemented")
}
func (DockerManagerServer) PruneContainers(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PruneContainers not implemented")
}
func (DockerManagerServer) CreateImage(context.Context, *docker_manager.CreateImageRequest) (*docker_manager.CreateImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImage not implemented")
}
func (DockerManagerServer) DeleteImage(context.Context, *docker_manager.DeleteImageRequest) (*docker_manager.DeleteImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (DockerManagerServer) GetImage(context.Context, *docker_manager.GetImageRequest) (*docker_manager.GetImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (DockerManagerServer) PruneImages(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PruneImages not implemented")
}
