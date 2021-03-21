package rpc

import (
	"context"
	containerManager "docker_manager/proto/container_server"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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

func PullImage(ctx context.Context, repository, tag string, authConf *containerManager.AuthConfig) (imageAttr *containerManager.ImageAttr, err error) {
	req := &containerManager.PullImage_Request{
		Repository: repository,
		Tag:        tag,
		AuthConfig: authConf,
	}
	resp, err := containerManagerClient.PullImage(ctx, req)
	if err != nil {
		return
	}

	return resp.ImageAttr, nil
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

func UploadImage(ctx context.Context, imageURL string) (imageAttr []*containerManager.ImageAttr, err error) {
	stream, err := containerManagerClient.LoadImage(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := http.Get(imageURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer res.Body.Close()

	image, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	logrus.Info(len(image))
	const size = 1 << 20
	n := len(image)/(size) + func() int {
		if len(image)&(size) > 0 {
			return 1
		}
		return 0
	}()
	for i := 0; i < n-1; i++ {
		req := &containerManager.LoadImage_Request{
			Data: image[i*size : i*size+size],
		}
		err = stream.Send(req)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	req := &containerManager.LoadImage_Request{
		Data: image[(n-1)*size:],
	}
	err = stream.Send(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return resp.ImageAttr, nil
}
