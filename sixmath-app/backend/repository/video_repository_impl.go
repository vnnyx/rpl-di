package repository

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/helper"
	"rpl-sixmath/model"

	"gorm.io/gorm"
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
	err = repo.DB.Where("video_id", videoId).First(&response).Error
	return response, err
}

func (repo *VideoRepositoryImpl) FindOneRandomVideo() (response entity.Video, err error) {
	err = repo.DB.Raw("SELECT * from videos order by rand()").Take(&response).Scan(&response).Error
	return response, err
}

func (repo *VideoRepositoryImpl) FindRecommendedVideo(pagination model.Pagination) *model.Pagination {
	var video []*entity.Video
	repo.DB.Scopes(helper.Paginate(video, &pagination, repo.DB)).Where("video_id != ?", pagination.MainVideoId).Find(&video)
	pagination.Rows = video

	return &pagination
}

func (repo *VideoRepositoryImpl) FindAllVideo(orderBy string) (response []entity.Video, err error) {
	if orderBy == "alphabet" {
		err = repo.DB.Order("title").Find(&response).Error
	} else {
		err = repo.DB.Order("created_at desc").Find(&response).Error
	}
	return response, err
}
