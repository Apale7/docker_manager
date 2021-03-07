package model

import (
	"docker_manager/proto/docker_manager"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func (i *Image) ToGRPCImage() *docker_manager.Image {
	return &docker_manager.Image{
		Id:       i.ImageID,
		Author:   i.Author,
		RepoTags: getTags(i.RepoTags),
		Created:  i.CreatedAt.Local().Unix(),
		Size:     i.ImageSize,
	}
}

// db中RepoTags字段存储的是json数组字符串，需要解析成数组返回
func getTags(jsonStr string) (res []string) {
	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		logrus.Warnf("[getTags] json Unmarshal error: %v", err)
	}

	return res
}
