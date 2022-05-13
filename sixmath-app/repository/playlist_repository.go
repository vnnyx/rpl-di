package repository

import (
	"rpl-sixmath/entity"
)

type PlaylistRepository interface {
	InsertPlaylist(playlist entity.Playlist) (entity.Playlist, error)
	UpdatePlaylist(playlist entity.Playlist) (entity.Playlist, error)
	DeletePlaylist(playlistId int) error
	FindPlaylistById(playlistId int) (response entity.Playlist, err error)
	FindPlaylistAll() (response []entity.Playlist, err error)
}
