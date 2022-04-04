package repository

import (
	"gorm.io/gorm"
	"rpl-sixmath/entity"
)

type VideoRepositoryImpl struct {
	DB *gorm.DB
}

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
