package repository

import (
	"gorm.io/gorm"
	"rpl-sixmath/entity"
	"rpl-sixmath/helper"
	"rpl-sixmath/model"
)

type VideoRepositoryImpl struct {
	DB *gorm.DB
}

func NewVideoRepository(DB *gorm.DB) VideoRepository {
	return &VideoRepositoryImpl{DB: DB}
}

func (repo *VideoRepositoryImpl) InsertVideo(video entity.Video) (entity.Video, error) {
	err := repo.DB.Create(&video).Error
	return video, err
}

func (repo *VideoRepositoryImpl) UpdateVideo(video entity.Video) (entity.Video, error) {
	err := repo.DB.Where("video_id", video.VideoId).Updates(&video).Error
	return video, err
}

func (repo *VideoRepositoryImpl) DeleteVideo(videoId int) error {
	err := repo.DB.Where("video_id", videoId).Delete(&entity.Video{}).Error
	return err
}

func (repo *VideoRepositoryImpl) FindVideoById(videoId int) (response entity.Video, err error) {
	err = repo.DB.Where("video_id", videoId).Find(&response).Error
	return response, err
}

func (repo *VideoRepositoryImpl) FindOneRandomVideo() (response entity.Video, err error) {
	err = repo.DB.Raw("SELECT * from videos order by rand()").Take(&response).Scan(&response).Error
	return response, err
}

func (repo *VideoRepositoryImpl) FindAllVideo(pagination model.Pagination) *model.Pagination {
	var video []*entity.Video
	repo.DB.Scopes(helper.Paginate(video, &pagination, repo.DB)).Where("video_id != ?", pagination.MainVideoId).Find(&video)
	pagination.Rows = video

	return &pagination
}
