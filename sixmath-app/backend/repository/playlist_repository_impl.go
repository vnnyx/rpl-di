package repository

import (
	"gorm.io/gorm"
	"rpl-sixmath/entity"
)

type PlaylistRepositoryImpl struct {
	DB *gorm.DB
}

func NewPlaylistRepository(DB *gorm.DB) PlaylistRepository {
	return &PlaylistRepositoryImpl{DB: DB}
}

func (repo *PlaylistRepositoryImpl) InsertPlaylist(playlist entity.Playlist) (entity.Playlist, error) {
	err := repo.DB.Table("playlists").Create(&playlist).Error
	return playlist, err
}

func (repo *PlaylistRepositoryImpl) UpdatePlaylist(playlist entity.Playlist) (entity.Playlist, error) {
	err := repo.DB.Table("playlists").Where("playlist_id", playlist.PlaylistId).Updates(&playlist).Error
	return playlist, err
}

func (repo *PlaylistRepositoryImpl) DeletePlaylist(playlistId int) error {
	err := repo.DB.Table("playlists").Where("playlist_id", playlistId).Delete(&entity.Playlist{}).Error
	return err
}

func (repo *PlaylistRepositoryImpl) FindPlaylistById(playlistId int) (response entity.Playlist, err error) {
	err = repo.DB.Table("playlists").Where("playlist_id", playlistId).Find(&response).Error
	return response, err
}

func (repo *PlaylistRepositoryImpl) FindPlaylistAll() (response []entity.Playlist, err error) {
	err = repo.DB.Table("playlists").Find(&response).Error
	return response, err
}
