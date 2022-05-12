package service

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
	"rpl-sixmath/repository"
	"rpl-sixmath/validation"
)

type VideoServiceImpl struct {
	VideoRepository repository.VideoRepository
}

func NewVideoService(videoRepository *repository.VideoRepository) VideoService {
	return &VideoServiceImpl{VideoRepository: *videoRepository}
}

func (service *VideoServiceImpl) CreateVideo(request model.VideoCreateRequest) (response model.VideoResponse) {
	validation.VideoValidate(request)

	video := entity.VideoEntity{
		PlaylistId: request.PlaylistId,
		Title:      request.Title,
		URLVideo:   request.URLVideo,
		Deskripsi:  request.Deskripsi,
	}

	_, _ = service.VideoRepository.InsertVideo(video)
	response = model.VideoResponse{
		VideoId:    video.VideoId,
		PlaylistId: video.PlaylistId,
		Title:      video.Title,
		URLVideo:   video.URLVideo,
		Deskripsi:  video.Deskripsi,
	}

	return response
}

func (service *VideoServiceImpl) UpdateVideo(request model.VideoUpdateRequest) (response model.VideoResponse) {
	validation.VideoUpdateValidate(request)

	video, _ := service.VideoRepository.FindVideoById(request.VideoId)
	video = entity.VideoEntity{
		VideoId:    request.VideoId,
		PlaylistId: request.PlaylistId,
		Title:      request.Title,
		URLVideo:   request.URLVideo,
		Deskripsi:  request.Deskripsi,
	}

	_, _ = service.VideoRepository.UpdateVideo(video)
	response = model.VideoResponse{
		VideoId:    video.VideoId,
		PlaylistId: video.PlaylistId,
		Title:      video.Title,
		URLVideo:   video.URLVideo,
		Deskripsi:  video.Deskripsi,
	}

	return response
}

func (service *VideoServiceImpl) GetMainVideo() (response model.VideoResponse) {
	video, _ := service.VideoRepository.FindOneRandomVideo()
	response = model.VideoResponse{
		VideoId:    video.VideoId,
		PlaylistId: video.PlaylistId,
		Title:      video.Title,
		URLVideo:   video.URLVideo,
		Deskripsi:  video.Deskripsi,
	}
	return response
}

func (service *VideoServiceImpl) DeleteVideo(videoId int) {
	_ = service.VideoRepository.DeleteVideo(videoId)
}

func (service *VideoServiceImpl) GetRecommendedVideo() (response []model.VideoResponse) {
	videos, _ := service.VideoRepository.FindAllVideo()
	for _, video := range videos {
		response = append(response, model.VideoResponse{
			VideoId:    video.VideoId,
			PlaylistId: video.PlaylistId,
			Title:      video.Title,
			URLVideo:   video.URLVideo,
			Deskripsi:  video.Deskripsi,
		})
	}
	return response
}
