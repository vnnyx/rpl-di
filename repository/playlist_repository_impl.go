package repository

import (
	"rpl-sixmath/entity"

	"gorm.io/gorm"
)

type playListRepositoryImpl struct {
	DB *gorm.DB
}

func NewPlayListRepository(DB *gorm.DB) PlayListRepository {
	return playListRepositoryImpl{DB: DB}
}

func (repo playListRepositoryImpl) CreatePlayList(Playlist entity.PlaylistEntity) (response entity.PlaylistEntity, err error) {
	err = repo.DB.Create(&Playlist).Error
	return Playlist, err
}

func (repo playListRepositoryImpl) DeletePlayList(playlistID int) (response entity.PlaylistEntity, err error) {
	err = repo.DB.Where("playlist_id", playlistID).Delete(&response).Error
	return response, err
}

func (repo playListRepositoryImpl) UpdatePlayList(Playlist entity.PlaylistEntity) (response entity.PlaylistEntity, err error) {
	err = repo.DB.Where("playlist_id", Playlist.Id).Updates(&response).Error
	return response, err
}

func (repo playListRepositoryImpl) GetPlayListByTeacher(teacherID int) (response []entity.PlaylistEntity, err error) {
	err = repo.DB.Where("teacher_id", teacherID).Find(&response).Error
	return response, err
}

func (repo playListRepositoryImpl) GetPlayListByID(playlistID int) (response entity.PlaylistEntity, err error) {
	err = repo.DB.Where("playlist_id", playlistID).First(&response).Error
	return response, err
}
