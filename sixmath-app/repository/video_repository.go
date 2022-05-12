package repository

import "rpl-sixmath/entity"

type VideoRepository interface {
	InsertVideo(video entity.VideoEntity) (entity.VideoEntity, error)
	UpdateVideo(video entity.VideoEntity) (entity.VideoEntity, error)
	DeleteVideo(videoId int) error
	FindVideoById(videoId int) (response entity.VideoEntity, err error)
	FindOneRandomVideo() (response entity.VideoEntity, err error)
	FindAllVideo() (response []entity.VideoEntity, err error)
}
