package repository

import "rpl-sixmath/entity"

type PlayListRepository interface {
	CreatePlayList(Playlist entity.PlaylistEntity) (response entity.PlaylistEntity, err error)
	DeletePlayList(playlistID int) (response entity.PlaylistEntity, err error)
	UpdatePlayList(Playlist entity.PlaylistEntity) (response entity.PlaylistEntity, err error)
	GetPlayListByTeacher(teacherID int) (response []entity.PlaylistEntity, err error)
	GetPlayListByID(playlistID int) (response entity.PlaylistEntity, err error)
}
