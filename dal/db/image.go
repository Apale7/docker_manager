package db

import (
	"context"
	"docker_manager/dal/db/model"
	"errors"

	oe "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 有imageIDs就直接按id查，没有就按userID在user_image表中查imageIDs
func GetImage(ctx context.Context, userIDs []uint32, imageIDs []string) (images []*model.Image, err error) {
	db := db.WithContext(ctx)
	if len(imageIDs) > 0 {
		err = db.Model(&model.Image{}).Where("image_id in ?", imageIDs).Find(&images).Error
	} else {
		err = db.Model(&model.UserImage{}).Select("image_id").Where("user_id in ?", userIDs).Pluck("image_id", &imageIDs).Error
		if err != nil {
			return nil, oe.WithStack(err)
		}
		err = db.Model(&model.Image{}).Where("image_id in ?", imageIDs).Find(&images).Error

	}
	err = oe.WithStack(err)

	return
}

func CreateImage(ctx context.Context, userID uint32, image *model.Image) (err error) {
	db := db.WithContext(ctx)
	err = db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.FirstOrCreate(image, *image).Error
		if err != nil {
			logrus.Warnf("[CreateImage] create image error, err: %v", err)
			return
		}

		if image.ID <= 0 {
			panic("invalid image_id")
		}

		err = tx.Create(&model.UserImage{UserID: userID, ImageID: image.ImageID}).Error
		return
	})
	return
}

func DeleteImage(ctx context.Context, userID uint32, imageID string) (err error) {
	userImage := model.UserImage{
		UserID:  userID,
		ImageID: imageID,
	}
	db := db.WithContext(ctx)
	db = db.Unscoped().WithContext(ctx).Where("user_id = ? AND image_id = ?", userID, imageID).Delete(&userImage)
	if db.Error != nil {
		return
	}
	if db.RowsAffected <= 0 {
		return errors.New("delete nothing")
	}
	return
}
