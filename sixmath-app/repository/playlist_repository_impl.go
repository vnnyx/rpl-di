package repository

import (
	"rpl-sixmath/entity"

	"gorm.io/gorm"
)

type PlaylistRepositoryImpl struct {
	DB *gorm.DB
}

func NewPlayListRepository(DB *gorm.DB) PlaylistRepository {
	return &PlaylistRepositoryImpl{DB: DB}
}

func (repo PlaylistRepositoryImpl) CreatePlayList(Playlist entity.PlaylistEntity) (response entity.PlaylistEntity, err error) {
	err = repo.DB.Create(&Playlist).Error
	return Playlist, err
}

func (repo PlaylistRepositoryImpl) DeletePlayList(playlistID int) (response entity.PlaylistEntity, err error) {
	err = repo.DB.Where("playlist_id", playlistID).Delete(&response).Error
	return response, err
}

func (repo PlaylistRepositoryImpl) UpdatePlayList(Playlist entity.PlaylistEntity) (response entity.PlaylistEntity, err error) {
	err = repo.DB.Where("playlist_id", Playlist.Id).Updates(&response).Error
	return response, err
}

func (repo PlaylistRepositoryImpl) GetPlayListByTeacher(teacherID int) (response []entity.PlaylistEntity, err error) {
	err = repo.DB.Where("teacher_id", teacherID).Find(&response).Error
	return response, err
}

func (repo PlaylistRepositoryImpl) GetPlayListByID(playlistID int) (response entity.PlaylistEntity, err error) {
	err = repo.DB.Where("playlist_id", playlistID).First(&response).Error
	return response, err
}
