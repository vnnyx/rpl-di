package service

import (
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
	"rpl-sixmath/repository"
)

type PlaylistServiceImpl struct {
	PlaylistRepository repository.PlaylistRepository
}

func NewPlaylistService(playlistRepository *repository.PlaylistRepository) PlaylistService {
	return &PlaylistServiceImpl{PlaylistRepository: *playlistRepository}
}

func (service *PlaylistServiceImpl) CreatePlaylist(request model.PlaylistCreateRequest) (response model.PlaylistCreateResponse) {
	playlist := entity.PlaylistEntity{
		Title:     request.Title,
		TeacherId: request.TeacherId,
	}
	_, _ = service.PlaylistRepository.InsertPlaylist(playlist)
	response = model.PlaylistCreateResponse{
		Title:     playlist.Title,
		TeacherId: playlist.TeacherId,
	}
	return response
}
