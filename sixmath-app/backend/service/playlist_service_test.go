package service

import (
	"github.com/stretchr/testify/assert"
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
	"rpl-sixmath/repository/mocks"
	"testing"
)

var (
	playlistEntity = entity.Playlist{
		PlaylistId: 1,
		Title:      "PlaylistTest",
		TeacherId:  1,
	}
)

func TestPlaylistServiceImpl_CreatePlaylistSuccess(t *testing.T) {
	playlistRepoMock := new(mocks.PlaylistRepository)

	playlist := entity.Playlist{
		Title:     playlistEntity.Title,
		TeacherId: playlistEntity.TeacherId,
	}

	playlistRepoMock.On("InsertPlaylist", playlist).Return(playlist, nil)

	playlistService := NewPlaylistService(playlistRepoMock)

	request := model.PlaylistCreateRequest{
		Title:     playlistEntity.Title,
		TeacherId: playlistEntity.TeacherId,
	}

	response, err := playlistService.CreatePlaylist(request)
	assert.Nil(t, err)
	assert.Equal(t, playlistEntity.Title, response.Title)
	assert.Equal(t, playlistEntity.TeacherId, response.TeacherId)
}
