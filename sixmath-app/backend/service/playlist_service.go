package service

import "rpl-sixmath/model"

type PlaylistService interface {
	CreatePlaylist(request model.PlaylistCreateRequest) (response model.PlaylistCreateResponse, err error)
}
