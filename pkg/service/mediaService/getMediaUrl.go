package mediaService

import (
	"go-admin/pkg/models"
	"gorm.io/gorm"
)

func GetMediaUrlArr(db *gorm.DB, mediaId []int64) ([]string, error) {
	mediaMap, err := GetMediaUrlMap(db, mediaId)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for k := range mediaMap {
		url := mediaMap[k]
		result = append(result, url)
	}
	return result, nil
}

func GetMediaUrlMap(db *gorm.DB, mediaIds []int64) (map[int64]string, error) {
	medias, err := GetMediasArr(db, mediaIds)
	if err != nil {
		return nil, err
	}

	result := make(map[int64]string)
	for i := range medias {
		media := medias[i]
		result[media.ID] = media.OriginalURL
	}
	return result, nil
}

func GetMediasArr(db *gorm.DB, mediaIds []int64) ([]*models.Media, error) {
	mediaMap, err := GetMediaMap(db, mediaIds)
	if err != nil {
		return nil, err
	}
	result := make([]*models.Media, 0)
	for i := range mediaIds {
		result = append(result, mediaMap[mediaIds[i]])
	}
	return result, nil
}

func GetMediaMap(db *gorm.DB, mediaIds []int64) (map[int64]*models.Media, error) {
	medias := make([]*models.Media, 0)
	err := db.Where("id in (?)", mediaIds).Find(&medias).Error
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*models.Media)
	for i := range medias {
		m := medias[i]
		result[m.ID] = m
	}
	return result, nil
}
