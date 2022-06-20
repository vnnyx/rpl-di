package service

import (
	"errors"
	"rpl-sixmath/entity"
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
	"rpl-sixmath/repository"
	"rpl-sixmath/validation"
)

type VideoServiceImpl struct {
	VideoRepository repository.VideoRepository
}

func NewVideoService(videoRepository repository.VideoRepository) VideoService {
	return &VideoServiceImpl{VideoRepository: videoRepository}
}

func (service *VideoServiceImpl) CreateVideo(request model.VideoCreateRequest) (response model.VideoResponse, err error) {
	validation.VideoValidate(request)

	video := entity.Video{
		TeacherUsername: request.Teacher,
		Title:           request.Title,
		URLVideo:        request.URLVideo,
		Description:     request.Description,
	}

	video, err = service.VideoRepository.InsertVideo(video)
	if err != nil {
		return model.VideoResponse{}, err
	}
	response = model.VideoResponse{
		VideoId:     video.VideoId,
		Teacher:     video.TeacherUsername,
		Title:       video.Title,
		URLVideo:    video.URLVideo,
		Description: video.Description,
	}

	return response, nil
}

func (service *VideoServiceImpl) UpdateVideo(request model.VideoUpdateRequest) (response model.VideoResponse, err error) {
	validation.VideoUpdateValidate(request)
	video, err := service.VideoRepository.FindVideoById(request.VideoId)
	exception.PanicIfNeeded(err)
	if (video == entity.Video{}) {
		return model.VideoResponse{}, errors.New("VIDEO_NOT_FOUND")
	}
	video = entity.Video{
		VideoId:         request.VideoId,
		TeacherUsername: request.Teacher,
		Title:           request.Title,
		URLVideo:        request.URLVideo,
		Description:     request.Description,
	}

	_, err = service.VideoRepository.UpdateVideo(video)
	if err != nil {
		return model.VideoResponse{}, err
	}
	response = model.VideoResponse{
		VideoId:     video.VideoId,
		Teacher:     video.TeacherUsername,
		Title:       video.Title,
		URLVideo:    video.URLVideo,
		Description: video.Description,
	}

	return response, nil
}

func (service *VideoServiceImpl) DetailVideo(videoId int) (response model.VideoResponse, err error) {
	video, err := service.VideoRepository.FindVideoById(videoId)
	if err != nil {
		return model.VideoResponse{}, errors.New("VIDEO_NOT_FOUND")
	}
	response = model.VideoResponse{
		VideoId:     video.VideoId,
		Teacher:     video.TeacherUsername,
		Title:       video.Title,
		URLVideo:    video.URLVideo,
		Description: video.Description,
	}
	return response, nil
}

func (service *VideoServiceImpl) DeleteVideo(videoId int) error {
	video, err := service.VideoRepository.FindVideoById(videoId)
	exception.PanicIfNeeded(err)
	if (video == entity.Video{}) {
		return errors.New("VIDEO_NOT_FOUND")
	}
	_ = service.VideoRepository.DeleteVideo(videoId)
	return nil
}

func (service *VideoServiceImpl) GetRecommendedVideo(pagination model.Pagination) (response model.Pagination) {
	video := service.VideoRepository.FindRecommendedVideo(pagination)
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

func (service *VideoServiceImpl) GetAllVideo(orderBy string) (response []model.VideoResponse, err error) {
	videos, err := service.VideoRepository.FindAllVideo(orderBy)
	if err != nil {
		return response, err
	}
	for _, video := range videos {
		response = append(response, model.VideoResponse{
			VideoId:     video.VideoId,
			Teacher:     video.TeacherUsername,
			Title:       video.Title,
			URLVideo:    video.URLVideo,
			Description: video.Description,
		})
	}
	return response, nil
}
