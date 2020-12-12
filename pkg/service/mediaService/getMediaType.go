package mediaService

import (
	"go-admin/pkg/models"
	"gorm.io/gorm"
)

func GetMediaType(db *gorm.DB, tType int64) (*models.MediaType, error) {
	result := &models.MediaType{}
	return result, db.Where("id=?", tType).First(result).Error
}
