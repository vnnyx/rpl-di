package service

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
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

	video := entity.Video{
		PlaylistId: request.PlaylistId,
		Title:      request.Title,
		URLVideo:   request.URLVideo,
		Deskripsi:  request.Deskripsi,
	}

	video, err := service.VideoRepository.InsertVideo(video)
	if err != nil {
		exception.PanicIfNeeded("PLAYLIST_NOT_FOUND")
	}
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
	video = entity.Video{
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

func (service *VideoServiceImpl) GetRecommendedVideo(pagination model.Pagination) (response model.Pagination) {
	video := service.VideoRepository.FindAllVideo(pagination)
	response = model.Pagination{
		Limit:      video.Limit,
		Page:       video.Page,
		Sort:       video.Sort,
		TotalRows:  video.TotalRows,
		TotalPages: video.TotalPages,
		Rows:       video.Rows,
	}
	return response
}
