package service

import "rpl-sixmath/model"

type VideoService interface {
	CreateVideo(request model.VideoCreateRequest) (response model.VideoResponse, err error)
	UpdateVideo(request model.VideoUpdateRequest) (response model.VideoResponse, err error)
	DetailVideo(videoId int) (response model.VideoResponse, err error)
	DeleteVideo(videoId int) error
	GetRecommendedVideo(pagination model.Pagination) (response model.Pagination)
	GetAllVideo(orderBy string) (response []model.VideoResponse, err error)
}
