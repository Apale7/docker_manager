package db

import (
	"context"
	"docker_manager/dal/db/model"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateContainer(ctx context.Context, userID uint32, container *model.Container) (err error) {
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&model.Container{}).Create(&container).Error
		if err != nil {
			logrus.Warnf("[CreateContainer] create container error, err: %v", err)
			return
		}

		if container.ID <= 0 {
			panic("invalid container_id")
		}

		userContainer := model.UserContainer{
			UserID:      userID,
			ContainerID: container.ContainerID,
		}
		err = tx.Model(&model.UserContainer{}).Create(&userContainer).Error
		if err != nil {
			logrus.Warnf("[CreateContainer] create user_container error, err: %v", err)
			return
		}
		return nil
	})

	return
}

func DeleteContainer(ctx context.Context, userID uint32, containerID string) (err error) {
	userContainer := model.UserContainer{
		UserID:      userID,
		ContainerID: containerID,
	}

	db := db.Unscoped().Model(&model.UserContainer{}).Where("user_id = ? AND container_id = ?", userID, containerID).Delete(&userContainer)
	if db.Error != nil {
		logrus.Warnf("DeleteContainer error: %v", err)
		return
	}
	if db.RowsAffected <= 0 {
		logrus.Warnln("delete nothing")
		return errors.New("delete nothing")
	}
	return
}

func GetContainer(ctx context.Context, userID uint32, containerID string) (containers []*model.Container, err error) {
	db := db.WithContext(ctx)
	if containerID != "" {
		err = db.Model(&model.Container{}).Where("container_id = ?", containerID).Find(&containers).Error
		return
	}
	var containerIDs []string
	err = db.Model(&model.UserContainer{}).Where("user_id = ?", userID).Pluck("container_id", &containerIDs).Error
	if err != nil {
		logrus.Warnf("Get ContainerIDs error: %v", err)
		return
	}
	err = db.Model(&model.Container{}).Where("container_id in (?)", containerIDs).Find(&containers).Error
	return
}
