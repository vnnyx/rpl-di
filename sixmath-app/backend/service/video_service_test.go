package service

import (
	"github.com/stretchr/testify/assert"
	"rpl-sixmath/entity"
	"rpl-sixmath/model"
	"rpl-sixmath/repository/mocks"
	"testing"
)

var (
	videoEntity = entity.Video{
		PlaylistId: 1,
		Title:      "title",
		URLVideo:   "urlvideo",
		Deskripsi:  "deskripsi",
	}
)

func TestCreateVideoSuccess(t *testing.T) {
	videoRepoMock := new(mocks.VideoRepository)

	video := entity.Video{
		PlaylistId: videoEntity.PlaylistId,
		Title:      videoEntity.Title,
		URLVideo:   videoEntity.URLVideo,
		Deskripsi:  videoEntity.Deskripsi,
	}

	videoRepoMock.On("InsertVideo", video).Return(video, nil)

	videoService := NewVideoService(videoRepoMock)

	request := model.VideoCreateRequest{
		PlaylistId: videoEntity.PlaylistId,
		Title:      videoEntity.Title,
		URLVideo:   videoEntity.URLVideo,
		Deskripsi:  videoEntity.Deskripsi,
	}

	response, err := videoService.CreateVideo(request)
	assert.Nil(t, err)
	assert.Equal(t, videoEntity.PlaylistId, response.PlaylistId)
	assert.Equal(t, videoEntity.URLVideo, response.URLVideo)
}
