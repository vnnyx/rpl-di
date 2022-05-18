package service

import "rpl-sixmath/model"

type VideoService interface {
	CreateVideo(request model.VideoCreateRequest) (response model.VideoResponse)
	UpdateVideo(request model.VideoUpdateRequest) (response model.VideoResponse)
	GetMainVideo() (response model.VideoResponse)
	DeleteVideo(videoId int)
	GetRecommendedVideo(pagination model.Pagination) (response model.Pagination)
}
