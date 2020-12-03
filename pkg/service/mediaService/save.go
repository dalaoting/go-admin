package mediaService

import (
	"errors"
	"go-admin/pkg/models"
	"gorm.io/gorm"
)

func Save(db *gorm.DB, media *models.Media) error {
	if media == nil {
		return errors.New("媒体为空")
	}
	return db.Create(&media).Error
}
