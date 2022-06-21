package repository

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
)

type VideoRepository interface {
	InsertVideo(video entity.Video) (entity.Video, error)
	UpdateVideo(video entity.Video) (entity.Video, error)
	DeleteVideo(videoId int) error
	FindVideoById(videoId int) (response entity.Video, err error)
	FindOneRandomVideo() (response entity.Video, err error)
	FindRecommendedVideo(pagination model.Pagination) *model.Pagination
	FindAllVideo(orderBy string) (response []entity.Video, err error)
}
