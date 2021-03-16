package db

import (
	"context"
	"docker_manager/dal/db/model"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetImage(ctx context.Context, userID uint32, imageID string) (images []*model.Image, err error) {
	db := db.WithContext(ctx)
	if imageID != "" {
		err = db.Where("image_id=?", imageID).Find(&images).Error
		return
	}
	var imageIDs []string
	err = db.Where("user_id=?", userID).Pluck("image_id", imageIDs).Error
	if err != nil || len(imageIDs) == 0 {
		return
	}
	err = db.Where("image_id in ?", imageIDs).Find(&images).Error
	return
}

func CreateImage(ctx context.Context, userID uint32, image *model.Image) (err error) {
	err = db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Create(image).Error
		if err != nil {
			logrus.Warnf("[CreateImage] create image error, err: %v", err)
			return
		}

		if image.ID <= 0 {
			panic("invalid image_id")
		}

		err = tx.FirstOrCreate(&model.UserImage{UserID: userID, ImageID: image.ImageID}).Error
		return
	})
	return
}

func DeleteImage(ctx context.Context, userID uint32, imageID string) (err error) {
	userImage := model.UserImage{
		UserID:	  userID,
		ImageID: imageID,
	}
	db := db.Unscoped().WithContext(ctx).Where("user_id = ? AND image_id = ?", userID, imageID).Delete(&userImage)
	if db.Error != nil{
		return
	}
	if db.RowsAffected <= 0 {
		return errors.New("delete nothing")
	}
	return
}