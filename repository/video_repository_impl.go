package repository

import (
	"rpl-sixmath/entity"

	"gorm.io/gorm"
)

type VideoRepositoryImpl struct {
	DB *gorm.DB
}

<<<<<<< HEAD
func NewVideoRepository(DB *gorm.DB) VideoRepository {
	return &VideoRepositoryImpl{DB: DB}
}

=======
func NewVideoRepositoryImpl(DB *gorm.DB) VideoRepository {
	return &VideoRepositoryImpl{DB: DB}
}
>>>>>>> 88382157a691e45b833da5899a3edb8b8fac3e78
func (repo VideoRepositoryImpl) AddVideo(video entity.VideoEntity) (response entity.VideoEntity, err error) {
	err = repo.DB.Create(&video).Error
	return video, err
}

func (repo VideoRepositoryImpl) UpdateVideo(video entity.VideoEntity) (response entity.VideoEntity, err error) {
	err = repo.DB.Where("video_id", video.Id).Updates(&video).Error
	return video, err
}

func (repo VideoRepositoryImpl) GetVideoById(videoId int) (response entity.VideoEntity, err error) {
	err = repo.DB.Where("video_id", videoId).First(&response).Error
	return response, err
}

func (repo VideoRepositoryImpl) GetVideosByPlaylistId(playListId int) (response []entity.VideoEntity, err error) {
	err = repo.DB.Where("playlist_id", playListId).Find(&response).Error
	return response, err
}

func (repo VideoRepositoryImpl) DeleteVideo(videoId int) error {
	err := repo.DB.Where("video_id", videoId).Delete(&entity.VideoEntity{}).Error
	return err
}
