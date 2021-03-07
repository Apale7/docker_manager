package Manage

import (
	"docker_manager/dal/db/model"
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func (i *ImageAttr) ToDBImage() *model.Image {
	tags, err := json.Marshal(&i.RepoTags)
	if err != nil {
		logrus.Warnf("[ToDBImage] json Marshal error: %v", err)
	}
	return &model.Image{
		ImageID:   i.Id,
		ImageSize: i.Size,
		Author:    i.Author,
		RepoTags:  string(tags),
	}
}
