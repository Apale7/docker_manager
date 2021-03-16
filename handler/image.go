package handler

import (
	"context"
	"docker_manager/dal/db"
	"docker_manager/dal/rpc"
	"docker_manager/dto"
	"docker_manager/proto/docker_manager"
)

func CreateImage(ctx context.Context, req *docker_manager.CreateImageRequest) (resp *docker_manager.CreateImageResponse, err error) {
	imageAttr, err := rpc.BuildImage(ctx, req.Dockerfile)
	if err != nil {
		return
	}

	err = db.CreateImage(ctx, req.UserId, dto.RPCImageToModelImage(imageAttr))
	if err != nil {
		return
	}
	resp = &docker_manager.CreateImageResponse{}
	return
}

func DeleteImage(ctx context.Context, req *docker_manager.DeleteImageRequest) (resp *docker_manager.DeleteImageResponse, err error) {
	err = rpc.DeleteImage(ctx, req.ImageId, req.Force)
	if err != nil {
		return
	}

	err = db.DeleteImage(ctx, req.UserId, req.ImageId)
	return
}

func GetImage(ctx context.Context, req *docker_manager.GetImageRequest) (resp *docker_manager.GetImageResponse, err error) {
	resp = &docker_manager.GetImageResponse{}
	images, err := db.GetImage(ctx, req.UserId, req.ImageId)

	if err != nil {
		return
	}

	resp.Images = make([]*docker_manager.Image, 0, len(images))
	for _, i := range images {
		resp.Images = append(resp.Images, dto.ModerImageToDockerManagerImage(i))
	}
	return
}
