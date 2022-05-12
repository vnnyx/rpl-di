package repository

import (
	"gorm.io/gorm"
	"rpl-sixmath/entity"
)

type VideoRepositoryImpl struct {
	DB *gorm.DB
}

func NewVideoRepository(DB *gorm.DB) VideoRepository {
	return &VideoRepositoryImpl{DB: DB}
}

func (repo *VideoRepositoryImpl) InsertVideo(video entity.VideoEntity) (entity.VideoEntity, error) {
	err := repo.DB.Create(&video).Error
	return video, err
}

func (repo *VideoRepositoryImpl) UpdateVideo(video entity.VideoEntity) (entity.VideoEntity, error) {
	err := repo.DB.Where("video_id", video.VideoId).Updates(&video).Error
	return video, err
}

func (repo *VideoRepositoryImpl) DeleteVideo(videoId int) error {
	err := repo.DB.Where("video_id", videoId).Delete(&entity.VideoEntity{}).Error
	return err
}

func (repo *VideoRepositoryImpl) FindVideoById(videoId int) (response entity.VideoEntity, err error) {
	err = repo.DB.Where("video_id", videoId).Find(&response).Error
	return response, err
}

func (repo *VideoRepositoryImpl) FindOneRandomVideo() (response entity.VideoEntity, err error) {
	err = repo.DB.Raw("SELECT * from video order by rand()").Take(&response).Scan(&response).Error
	return response, err
}

func (repo *VideoRepositoryImpl) FindAllVideo() (response []entity.VideoEntity, err error) {
	err = repo.DB.Find(&response).Error
	return response, err
}
