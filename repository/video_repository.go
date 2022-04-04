package repository

import "rpl-sixmath/entity"

type VideoRepository interface {
	AddVideo(video entity.VideoEntity) (response entity.VideoEntity, err error)
	UpdateVideo(video entity.VideoEntity) (response entity.VideoEntity, err error)
	GetVideoById(videoId int) (response entity.VideoEntity, err error)
	GetVideosByPlaylistId(playListId int) (response []entity.VideoEntity, err error)
	DeleteVideo(videoId int) error
}
