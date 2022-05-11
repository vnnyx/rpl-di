package repository

import "rpl-sixmath/entity"

type PlaylistRepository interface {
	InsertPlaylist(playlist entity.PlaylistEntity) (entity.PlaylistEntity, error)
	UpdatePlaylist(playlist entity.PlaylistEntity) (entity.PlaylistEntity, error)
	DeletePlaylist(playlistId int) error
	FindPlaylistById(playlistId int) (response entity.PlaylistEntity, err error)
	FindPlaylistAll() (response []entity.PlaylistEntity, err error)
}
