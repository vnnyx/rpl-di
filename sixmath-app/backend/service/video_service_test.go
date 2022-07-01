package service

import (
	"github.com/stretchr/testify/assert"
	"rpl-sixmath/entity"
	"rpl-sixmath/repository/mocks"
	"testing"
)

var (
	videoEntity = entity.Video{
		VideoId:         1,
		TeacherUsername: "testteacher",
		Title:           "testtitle",
		URLVideo:        "testurl",
		Description:     "testdeskripsi",
	}
)

func TestVideoServiceImpl_GetAllVideo(t *testing.T) {
	videoRepMock := new(mocks.VideoRepository)
	videoRepMock.On("FindAllVideo", "new").Return([]entity.Video{videoEntity}, nil)
	videoService := NewVideoService(videoRepMock)

	response, err := videoService.GetAllVideo("new")
	assert.Nil(t, err)
	assert.Equal(t, videoEntity.Title, response[0].Title)
}
