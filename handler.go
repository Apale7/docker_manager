package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"docker_manager/proto/docker_manager"
)

// DockerManagerServer DockerManagerServer impl
type DockerManagerServer struct {
}

// CreateContainer create a container for a user
func (DockerManagerServer) CreateContainer(ctx context.Context, req *docker_manager.CreateContainerRequest) (resp *docker_manager.CreateContainerResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContainer not implemented")
}

// DeleteContainer delete a container for a user
func (DockerManagerServer) DeleteContainer(ctx context.Context, req *docker_manager.DeleteContainerRequest) (resp *docker_manager.DeleteContainerResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContainer not implemented")
}

// GetContainer get containers by container_id or user_id
func (DockerManagerServer) GetContainer(ctx context.Context, req *docker_manager.GetContainerRequest) (resp *docker_manager.GetContainerResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainer not implemented")
}

// PruneContainers delete all unused containers
func (DockerManagerServer) PruneContainers(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PruneContainers not implemented")
}

// CreateImage create a image for a user
func (DockerManagerServer) CreateImage(ctx context.Context, req *docker_manager.CreateImageRequest) (resp *docker_manager.CreateImageResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImage not implemented")
}

// DeleteImage delete a image for a user
func (DockerManagerServer) DeleteImage(ctx context.Context, req *docker_manager.DeleteImageRequest) (resp *docker_manager.DeleteImageResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}

// GetImage get images by image_id or user_id
func (DockerManagerServer) GetImage(ctx context.Context, req *docker_manager.GetImageRequest) (resp *docker_manager.GetImageResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}

// PruneImages delete all unused images
func (DockerManagerServer) PruneImages(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PruneImages not implemented")
}

// GetAllContainers get all containers
func (DockerManagerServer) GetAllContainers(ctx context.Context, req *emptypb.Empty) (resp *docker_manager.GetAllContainersResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllContainers not implemented")
}

// GetAllImages get all images
func (DockerManagerServer) GetAllImages(ctx context.Context, req *emptypb.Empty) (resp *docker_manager.GetAllImagesResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllImages not implemented")
}
