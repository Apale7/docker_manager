package rpc

import (
	"context"
	containerManager "docker_manager/proto/container_server"
)

func GetAllImages(ctx context.Context) (images []*containerManager.ImageAttr, err error) {
	resp, err := containerManagerClient.ListImages(ctx, &containerManager.ListImages_Request{})
	if err != nil {
		return
	}
	return resp.Images, nil
}

func PruneImages(ctx context.Context) (err error) {
	_, err = containerManagerClient.PruneImages(ctx, &containerManager.PruneImages_Request{})

	if err != nil {
		return
	}
	return
}

func PullImage(ctx context.Context, repository, tag string, authConf *containerManager.AuthConfig) (err error) {
	req := &containerManager.PullImage_Request{
		Repository: repository,
		Tag:        tag,
		AuthConfig: authConf,
	}
	_, err = containerManagerClient.PullImage(ctx, req)
	if err != nil {
		return
	}

	return
}

func BuildImage(ctx context.Context, dockerfile []byte) (imageAttr *containerManager.ImageAttr, err error) {
	req := &containerManager.BuildImage_Request{Dockerfile: dockerfile}
	resp, err := containerManagerClient.BuildImage(ctx, req)
	if err != nil {
		return
	}

	return resp.ImageAttr, nil
}

func GetImage(ctx context.Context, imageID string) (image *containerManager.ImageAttr, err error) {
	req := &containerManager.GetImage_Request{ImageId: imageID}
	resp, err := containerManagerClient.GetImage(ctx, req)
	if err != nil {
		return
	}

	return resp.ImageAttr, nil
}

func DeleteImage(ctx context.Context, imageID string, force bool) (err error) {
	req := &containerManager.RemoveImage_Request{ImageId: imageID, Force: force}
	_, err = containerManagerClient.RemoveImage(ctx, req)
	if err != nil {
		return
	}

	return
}
